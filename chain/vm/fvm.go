package vm

import (
	"bytes"
	"context"

	"github.com/filecoin-project/lotus/chain/state"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/lotus/lib/sigs"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/blockstore"

	ffi "github.com/filecoin-project/filecoin-ffi"
	ffi_cgo "github.com/filecoin-project/filecoin-ffi/cgo"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var _ VMI = (*FVM)(nil)
var _ ffi_cgo.Externs = (*FvmExtern)(nil)

type FvmExtern struct {
	Rand
	blockstore.Blockstore
	epoch   abi.ChainEpoch
	lbState LookbackStateGetter
	base    cid.Cid
}

// Similar to the one in syscalls.go used by the Lotus VM, except it never errors
// Errors are logged and "no fault" is returned, which is functionally what go-actors does anyway
func (x *FvmExtern) VerifyConsensusFault(ctx context.Context, a, b, extra []byte) *ffi_cgo.ConsensusFaultWithGas {
	ret := &ffi_cgo.ConsensusFaultWithGas{
		// TODO: is this gonna be a problem on the receiving end? should we return address.NewIDAddress(0) instead?
		Target: address.Undef,
		Type:   ffi_cgo.ConsensusFaultNone,
	}

	// Note that block syntax is not validated. Any validly signed block will be accepted pursuant to the below conditions.
	// Whether or not it could ever have been accepted in a chain is not checked/does not matter here.
	// for that reason when checking block parent relationships, rather than instantiating a Tipset to do so
	// (which runs a syntactic check), we do it directly on the CIDs.

	// (0) cheap preliminary checks

	// can blocks be decoded properly?
	var blockA, blockB types.BlockHeader
	if decodeErr := blockA.UnmarshalCBOR(bytes.NewReader(a)); decodeErr != nil {
		log.Info("invalid consensus fault: cannot decode first block header: %w", decodeErr)
		return ret
	}

	if decodeErr := blockB.UnmarshalCBOR(bytes.NewReader(b)); decodeErr != nil {
		log.Info("invalid consensus fault: cannot decode second block header: %w", decodeErr)
		return ret
	}

	// Commented out from the Lotus VM code: FvmExtern only supports v14 and onwards
	// workaround chain halt
	//if build.IsNearUpgrade(blockA.Height, build.UpgradeOrangeHeight) {
	//	return nil, xerrors.Errorf("consensus reporting disabled around Upgrade Orange")
	//}
	//if build.IsNearUpgrade(blockB.Height, build.UpgradeOrangeHeight) {
	//	return nil, xerrors.Errorf("consensus reporting disabled around Upgrade Orange")
	//}

	// are blocks the same?
	if blockA.Cid().Equals(blockB.Cid()) {
		log.Info("invalid consensus fault: submitted blocks are the same")
		return ret
	}
	// (1) check conditions necessary to any consensus fault

	// were blocks mined by same miner?
	if blockA.Miner != blockB.Miner {
		log.Info("invalid consensus fault: blocks not mined by the same miner")
		return ret
	}

	ret.Target = blockA.Miner

	// block a must be earlier or equal to block b, epoch wise (ie at least as early in the chain).
	if blockB.Height < blockA.Height {
		log.Info("invalid consensus fault: first block must not be of higher height than second")
		return ret
	}

	ret.Epoch = blockB.Height

	faultType := ffi_cgo.ConsensusFaultNone

	// (2) check for the consensus faults themselves
	// (a) double-fork mining fault
	if blockA.Height == blockB.Height {
		faultType = ffi_cgo.ConsensusFaultDoubleForkMining
	}

	// (b) time-offset mining fault
	// strictly speaking no need to compare heights based on double fork mining check above,
	// but at same height this would be a different fault.
	if types.CidArrsEqual(blockA.Parents, blockB.Parents) && blockA.Height != blockB.Height {
		faultType = ffi_cgo.ConsensusFaultTimeOffsetMining
	}

	// (c) parent-grinding fault
	// Here extra is the "witness", a third block that shows the connection between A and B as
	// A's sibling and B's parent.
	// Specifically, since A is of lower height, it must be that B was mined omitting A from its tipset
	//
	//      B
	//      |
	//  [A, C]
	var blockC types.BlockHeader
	if len(extra) > 0 {
		if decodeErr := blockC.UnmarshalCBOR(bytes.NewReader(extra)); decodeErr != nil {
			log.Info("invalid consensus fault: cannot decode extra: %w", decodeErr)
			// just to match Lotus VM consensus, zero out any already-set faults
			faultType = ffi_cgo.ConsensusFaultNone
			return ret
		}

		if types.CidArrsEqual(blockA.Parents, blockC.Parents) && blockA.Height == blockC.Height &&
			types.CidArrsContains(blockB.Parents, blockC.Cid()) && !types.CidArrsContains(blockB.Parents, blockA.Cid()) {
			faultType = ffi_cgo.ConsensusFaultParentGrinding
		}
	}

	// (3) return if no consensus fault by now
	if faultType == ffi_cgo.ConsensusFaultNone {
		log.Info("invalid consensus fault: no fault detected")
		return ret
	}

	// else
	// (4) expensive final checks

	// check blocks are properly signed by their respective miner
	// note we do not need to check extra's: it is a parent to block b
	// which itself is signed, so it was willingly included by the miner
	gasUsed, sigErr := x.VerifyBlockSig(ctx, &blockA)
	ret.GasUsed += gasUsed
	if sigErr != nil {
		log.Info("invalid consensus fault: cannot verify first block sig: %w", sigErr)
		return ret
	}

	gasUsed, sigErr = x.VerifyBlockSig(ctx, &blockB)
	ret.GasUsed += gasUsed
	if sigErr != nil {
		log.Info("invalid consensus fault: cannot verify second block sig: %w", sigErr)
		return ret
	}

	ret.Type = faultType

	return ret
}

func (x *FvmExtern) VerifyBlockSig(ctx context.Context, blk *types.BlockHeader) (int64, error) {
	waddr, gasUsed, err := x.workerKeyAtLookback(ctx, blk.Miner, blk.Height)
	if err != nil {
		return gasUsed, err
	}

	return gasUsed, sigs.CheckBlockSignature(ctx, blk, waddr)
}

func (x *FvmExtern) workerKeyAtLookback(ctx context.Context, minerId address.Address, height abi.ChainEpoch) (address.Address, int64, error) {
	// Commented out from the Lotus VM code: FvmExtern only supports v14 and onwards
	//if x.networkVersion >= network.Version7 && height < x.epoch-policy.ChainFinality {
	//	return address.Undef, xerrors.Errorf("cannot get worker key (currEpoch %d, height %d)", ss.epoch, height)
	//}

	gasUsed := int64(0)
	gasAdder := func(gc GasCharge) {
		// technically not overflow safe, but that's fine
		gasUsed += gc.Total()
	}

	cstWithoutGas := cbor.NewCborStore(x.Blockstore)
	cbb := &gasChargingBlocks{gasAdder, PricelistByEpoch(x.epoch), x.Blockstore}
	cstWithGas := cbor.NewCborStore(cbb)

	lbState, err := x.lbState(ctx, height)
	if err != nil {
		return address.Undef, gasUsed, err
	}
	// get appropriate miner actor
	act, err := lbState.GetActor(minerId)
	if err != nil {
		return address.Undef, gasUsed, err
	}

	// use that to get the miner state
	mas, err := miner.Load(adt.WrapStore(ctx, cstWithGas), act)
	if err != nil {
		return address.Undef, gasUsed, err
	}

	info, err := mas.Info()
	if err != nil {
		return address.Undef, gasUsed, err
	}

	stateTree, err := state.LoadStateTree(cstWithoutGas, x.base)
	if err != nil {
		return address.Undef, gasUsed, err
	}

	raddr, err := ResolveToKeyAddr(stateTree, cstWithGas, info.Worker)
	if err != nil {
		return address.Undef, gasUsed, err
	}

	return raddr, gasUsed, nil
}

type FVM struct {
	fvm *ffi.FVM
}

func NewFVM(ctx context.Context, opts *VMOpts) (*FVM, error) {
	fvm, err := ffi.CreateFVM(0,
		&FvmExtern{Rand: opts.Rand, Blockstore: opts.Bstore, lbState: opts.LookbackState, base: opts.StateBase, epoch: opts.Epoch},
		opts.Epoch, opts.BaseFee, opts.FilVested, opts.NetworkVersion, opts.StateBase,
	)
	if err != nil {
		return nil, err
	}

	return &FVM{
		fvm: fvm,
	}, nil
}

func (vm *FVM) ApplyMessage(ctx context.Context, cmsg types.ChainMsg) (*ApplyRet, error) {
	msgBytes, err := cmsg.VMMessage().Serialize()
	if err != nil {
		return nil, xerrors.Errorf("serializing msg: %w", err)
	}

	ret, err := vm.fvm.ApplyMessage(msgBytes, uint(cmsg.ChainLength()))
	if err != nil {
		return nil, xerrors.Errorf("applying msg: %w", err)
	}

	return &ApplyRet{
		MessageReceipt: types.MessageReceipt{
			Return:   ret.Return,
			ExitCode: exitcode.ExitCode(ret.ExitCode),
			GasUsed:  ret.GasUsed,
		},
		GasCosts: &GasOutputs{
			// TODO: do the other optional fields eventually
			BaseFeeBurn:        abi.TokenAmount{},
			OverEstimationBurn: abi.TokenAmount{},
			MinerPenalty:       ret.MinerPenalty,
			MinerTip:           ret.MinerTip,
			Refund:             abi.TokenAmount{},
			GasRefund:          0,
			GasBurned:          0,
		},
		// TODO: do these eventually, not consensus critical
		ActorErr:       nil,
		ExecutionTrace: types.ExecutionTrace{},
		Duration:       0,
	}, nil
}

func (vm *FVM) ApplyImplicitMessage(ctx context.Context, cmsg *types.Message) (*ApplyRet, error) {
	msgBytes, err := cmsg.VMMessage().Serialize()
	if err != nil {
		return nil, xerrors.Errorf("serializing msg: %w", err)
	}

	ret, err := vm.fvm.ApplyImplicitMessage(msgBytes)
	if err != nil {
		return nil, xerrors.Errorf("applying msg: %w", err)
	}

	return &ApplyRet{
		MessageReceipt: types.MessageReceipt{
			Return:   ret.Return,
			ExitCode: exitcode.ExitCode(ret.ExitCode),
			GasUsed:  ret.GasUsed,
		},
		GasCosts: &GasOutputs{
			// TODO: do the other optional fields eventually
			BaseFeeBurn:        abi.TokenAmount{},
			OverEstimationBurn: abi.TokenAmount{},
			MinerPenalty:       ret.MinerPenalty,
			MinerTip:           ret.MinerTip,
			Refund:             abi.TokenAmount{},
			GasRefund:          0,
			GasBurned:          0,
		},
		// TODO: do these eventually, not consensus critical
		ActorErr:       nil,
		ExecutionTrace: types.ExecutionTrace{},
		Duration:       0,
	}, nil
}

func (vm *FVM) Flush(ctx context.Context) (cid.Cid, error) {
	return vm.fvm.Flush()
}
