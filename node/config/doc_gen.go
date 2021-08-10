// Code generated by github.com/filecoin-project/lotus/node/config/cfgdocgen. DO NOT EDIT.

package config

type DocField struct {
	Name    string
	Type    string
	Comment string
}

var Doc = map[string][]DocField{
	"API": []DocField{
		{
			Name: "ListenAddress",
			Type: "string",

			Comment: `Binding address for the Lotus API`,
		},
		{
			Name: "RemoteListenAddress",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "Timeout",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Backup": []DocField{
		{
			Name: "DisableMetadataLog",
			Type: "bool",

			Comment: `Note that in case of metadata corruption it might be much harder to recover
your node if metadata log is disabled`,
		},
	},
	"BatchFeeConfig": []DocField{
		{
			Name: "Base",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "PerSector",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"Chainstore": []DocField{
		{
			Name: "EnableSplitstore",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "Splitstore",
			Type: "Splitstore",

			Comment: ``,
		},
	},
	"Client": []DocField{
		{
			Name: "UseIpfs",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "IpfsOnlineMode",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "IpfsMAddr",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "IpfsUseForRetrieval",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "SimultaneousTransfers",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers`,
		},
	},
	"Common": []DocField{
		{
			Name: "API",
			Type: "API",

			Comment: ``,
		},
		{
			Name: "Backup",
			Type: "Backup",

			Comment: ``,
		},
		{
			Name: "Libp2p",
			Type: "Libp2p",

			Comment: ``,
		},
		{
			Name: "Pubsub",
			Type: "Pubsub",

			Comment: ``,
		},
	},
	"DAGStoreConfig": []DocField{
		{
			Name: "RootDir",
			Type: "string",

			Comment: `Path to the dagstore root directory. This directory contains three
subdirectories, which can be symlinked to alternative locations if
need be:
- ./transients: caches unsealed deals that have been fetched from the
storage subsystem for serving retrievals.
- ./indices: stores shard indices.
- ./datastore: holds the KV store tracking the state of every shard
known to the DAG store.
Default value: <LOTUS_MARKETS_PATH>/dagstore (split deployment) or
<LOTUS_MINER_PATH>/dagstore (monolith deployment)`,
		},
		{
			Name: "MaxConcurrentIndex",
			Type: "int",

			Comment: `The maximum amount of indexing jobs that can run simultaneously.
0 means unlimited.
Default value: 5.`,
		},
		{
			Name: "MaxConcurrentReadyFetches",
			Type: "int",

			Comment: `The maximum amount of unsealed deals that can be fetched simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Name: "MaxConcurrencyStorageCalls",
			Type: "int",

			Comment: `The maximum number of simultaneous inflight API calls to the storage
subsystem.
Default value: 100.`,
		},
		{
			Name: "GCInterval",
			Type: "Duration",

			Comment: `The time between calls to periodic dagstore GC, in time.Duration string
representation, e.g. 1m, 5m, 1h.
Default value: 1 minute.`,
		},
	},
	"DealmakingConfig": []DocField{
		{
			Name: "ConsiderOnlineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept online deals`,
		},
		{
			Name: "ConsiderOfflineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline deals`,
		},
		{
			Name: "ConsiderOnlineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept retrieval deals`,
		},
		{
			Name: "ConsiderOfflineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline retrieval deals`,
		},
		{
			Name: "ConsiderVerifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept verified deals`,
		},
		{
			Name: "ConsiderUnverifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept unverified deals`,
		},
		{
			Name: "PieceCidBlocklist",
			Type: "[]cid.Cid",

			Comment: `A list of Data CIDs to reject when making deals`,
		},
		{
			Name: "ExpectedSealDuration",
			Type: "Duration",

			Comment: `Maximum expected amount of time getting the deal into a sealed sector will take
This includes the time the deal will need to get transferred and published
before being assigned to a sector`,
		},
		{
			Name: "MaxDealStartDelay",
			Type: "Duration",

			Comment: `Maximum amount of time proposed deal StartEpoch can be in future`,
		},
		{
			Name: "PublishMsgPeriod",
			Type: "Duration",

			Comment: `When a deal is ready to publish, the amount of time to wait for more
deals to be ready to publish before publishing them all as a batch`,
		},
		{
			Name: "MaxDealsPerPublishMsg",
			Type: "uint64",

			Comment: `The maximum number of deals to include in a single PublishStorageDeals
message`,
		},
		{
			Name: "MaxProviderCollateralMultiplier",
			Type: "uint64",

			Comment: `The maximum collateral that the provider will put up against a deal,
as a multiplier of the minimum collateral bound`,
		},
		{
			Name: "SimultaneousTransfers",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers (storage+retrieval)`,
		},
		{
			Name: "Filter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of storage deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalFilter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of retrieval deals
see https://docs.filecoin.io/mine/lotus/miner-configuration/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalPricing",
			Type: "*RetrievalPricing",

			Comment: ``,
		},
	},
	"FeeConfig": []DocField{
		{
			Name: "DefaultMaxFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"FullNode": []DocField{
		{
			Name: "Client",
			Type: "Client",

			Comment: ``,
		},
		{
			Name: "Metrics",
			Type: "Metrics",

			Comment: ``,
		},
		{
			Name: "Wallet",
			Type: "Wallet",

			Comment: ``,
		},
		{
			Name: "Fees",
			Type: "FeeConfig",

			Comment: ``,
		},
		{
			Name: "Chainstore",
			Type: "Chainstore",

			Comment: ``,
		},
	},
	"Libp2p": []DocField{
		{
			Name: "ListenAddresses",
			Type: "[]string",

			Comment: `Binding address for the libp2p host - 0 means random port.
Format: multiaddress; see https://multiformats.io/multiaddr/`,
		},
		{
			Name: "AnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to explicitally announce to other peers. If not specified,
all interface addresses are announced
Format: multiaddress`,
		},
		{
			Name: "NoAnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to not announce
Format: multiaddress`,
		},
		{
			Name: "BootstrapPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "ProtectedPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "ConnMgrLow",
			Type: "uint",

			Comment: ``,
		},
		{
			Name: "ConnMgrHigh",
			Type: "uint",

			Comment: ``,
		},
		{
			Name: "ConnMgrGrace",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Metrics": []DocField{
		{
			Name: "Nickname",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "HeadNotifs",
			Type: "bool",

			Comment: ``,
		},
	},
	"MinerAddressConfig": []DocField{
		{
			Name: "PreCommitControl",
			Type: "[]string",

			Comment: `Addresses to send PreCommit messages from`,
		},
		{
			Name: "CommitControl",
			Type: "[]string",

			Comment: `Addresses to send Commit messages from`,
		},
		{
			Name: "TerminateControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DealPublishControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DisableOwnerFallback",
			Type: "bool",

			Comment: `DisableOwnerFallback disables usage of the owner address for messages
sent automatically`,
		},
		{
			Name: "DisableWorkerFallback",
			Type: "bool",

			Comment: `DisableWorkerFallback disables usage of the worker address for messages
sent automatically, if control addresses are configured.
A control address that doesn't have enough funds will still be chosen
over the worker address if this flag is set.`,
		},
	},
	"MinerFeeConfig": []DocField{
		{
			Name: "MaxPreCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxPreCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: `maxBatchFee = maxBase + maxPerSector * nSectors`,
		},
		{
			Name: "MaxCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: ``,
		},
		{
			Name: "MaxTerminateGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxWindowPoStGasFee",
			Type: "types.FIL",

			Comment: `WindowPoSt is a high-value operation, so the default fee should be high.`,
		},
		{
			Name: "MaxPublishDealsFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxMarketBalanceAddFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"MinerSubsystemConfig": []DocField{
		{
			Name: "EnableMining",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableSealing",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableSectorStorage",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableMarkets",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "SealerApiInfo",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "SectorIndexApiInfo",
			Type: "string",

			Comment: ``,
		},
	},
	"Pubsub": []DocField{
		{
			Name: "Bootstrapper",
			Type: "bool",

			Comment: `Run the node in bootstrap-node mode`,
		},
		{
			Name: "DirectPeers",
			Type: "[]string",

			Comment: `DirectPeers specifies peers with direct peering agreements. These peers are
connected outside of the mesh, with all (valid) message unconditionally
forwarded to them. The router will maintain open connections to these peers.
Note that the peering agreement should be reciprocal with direct peers
symmetrically configured at both ends.
Type: Array of multiaddress peerinfo strings, must include peerid (/p2p/12D3K...`,
		},
		{
			Name: "IPColocationWhitelist",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "RemoteTracer",
			Type: "string",

			Comment: ``,
		},
	},
	"RetrievalPricing": []DocField{
		{
			Name: "Strategy",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "Default",
			Type: "*RetrievalPricingDefault",

			Comment: ``,
		},
		{
			Name: "External",
			Type: "*RetrievalPricingExternal",

			Comment: ``,
		},
	},
	"RetrievalPricingDefault": []DocField{
		{
			Name: "VerifiedDealsFreeTransfer",
			Type: "bool",

			Comment: `VerifiedDealsFreeTransfer configures zero fees for data transfer for a retrieval deal
of a payloadCid that belongs to a verified storage deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "default".
default value is true`,
		},
	},
	"RetrievalPricingExternal": []DocField{
		{
			Name: "Path",
			Type: "string",

			Comment: `Path of the external script that will be run to price a retrieval deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "external".`,
		},
	},
	"SealingConfig": []DocField{
		{
			Name: "MaxWaitDealsSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be waiting for more deals to be packed in it before it begins sealing at any given time.
If the miner is accepting multiple deals in parallel, up to MaxWaitDealsSectors of new sectors will be created.
If more than MaxWaitDealsSectors deals are accepted in parallel, only MaxWaitDealsSectors deals will be processed in parallel
Note that setting this number too high in relation to deal ingestion rate may result in poor sector packing efficiency
0 = no limit`,
		},
		{
			Name: "MaxSealingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing at the same time when creating new CC sectors (0 = unlimited)`,
		},
		{
			Name: "MaxSealingSectorsForDeals",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing at the same time when creating new sectors with deals (0 = unlimited)`,
		},
		{
			Name: "WaitDealsDelay",
			Type: "Duration",

			Comment: `Period of time that a newly created sector will wait for more deals to be packed in to before it starts to seal.
Sectors which are fully filled will start sealing immediately`,
		},
		{
			Name: "AlwaysKeepUnsealedCopy",
			Type: "bool",

			Comment: `Whether to keep unsealed copies of deal data regardless of whether the client requested that. This lets the miner
avoid the relatively high cost of unsealing the data later, at the cost of more storage space`,
		},
		{
			Name: "FinalizeEarly",
			Type: "bool",

			Comment: `Run sector finalization before submitting sector proof to the chain`,
		},
		{
			Name: "CollateralFromMinerBalance",
			Type: "bool",

			Comment: `Whether to use available miner balance for sector collateral instead of sending it with each message`,
		},
		{
			Name: "AvailableBalanceBuffer",
			Type: "types.FIL",

			Comment: `Minimum available balance to keep in the miner actor before sending it with messages`,
		},
		{
			Name: "DisableCollateralFallback",
			Type: "bool",

			Comment: `Don't send collateral with messages even if there is no available balance in the miner actor`,
		},
		{
			Name: "BatchPreCommits",
			Type: "bool",

			Comment: `enable / disable precommit batching (takes effect after nv13)`,
		},
		{
			Name: "MaxPreCommitBatch",
			Type: "int",

			Comment: `maximum precommit batch size - batches will be sent immediately above this size`,
		},
		{
			Name: "PreCommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Name: "PreCommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deal in batch would start expiring`,
		},
		{
			Name: "AggregateCommits",
			Type: "bool",

			Comment: `enable / disable commit aggregation (takes effect after nv13)`,
		},
		{
			Name: "MinCommitBatch",
			Type: "int",

			Comment: `maximum batched commit size - batches will be sent immediately above this size`,
		},
		{
			Name: "MaxCommitBatch",
			Type: "int",

			Comment: ``,
		},
		{
			Name: "CommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Name: "CommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deals in batch would start expiring`,
		},
		{
			Name: "AggregateAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing commit aggregation, instead
submitting proofs to the chain individually`,
		},
		{
			Name: "TerminateBatchMax",
			Type: "uint64",

			Comment: ``,
		},
		{
			Name: "TerminateBatchMin",
			Type: "uint64",

			Comment: ``,
		},
		{
			Name: "TerminateBatchWait",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Splitstore": []DocField{
		{
			Name: "ColdStoreType",
			Type: "string",

			Comment: `ColdStoreType specifies the type of the coldstore.
It can be "universal" (default) or "discard" for discarding cold blocks.`,
		},
		{
			Name: "HotStoreType",
			Type: "string",

			Comment: `HotStoreType specifies the type of the hotstore.
Only currently supported value is "badger".`,
		},
		{
			Name: "MarkSetType",
			Type: "string",

			Comment: `MarkSetType specifies the type of the markset.
It can be "map" (default) for in memory marking or "badger" for on-disk marking.`,
		},
		{
			Name: "HotStoreMessageRetention",
			Type: "uint64",

			Comment: `HotStoreMessageRetention specifies the retention policy for messages, in finalities beyond
the compaction boundary; default is 0.`,
		},
		{
			Name: "HotStoreFullGCFrequency",
			Type: "uint64",

			Comment: `HotStoreFullGCFrequency specifies how often to perform a full (moving) GC on the hotstore.
A value of 0 disables, while a value 1 will do full GC in every compaction.
Default is 20 (about once a week).`,
		},
	},
	"StorageMiner": []DocField{
		{
			Name: "Subsystems",
			Type: "MinerSubsystemConfig",

			Comment: ``,
		},
		{
			Name: "Dealmaking",
			Type: "DealmakingConfig",

			Comment: ``,
		},
		{
			Name: "Sealing",
			Type: "SealingConfig",

			Comment: ``,
		},
		{
			Name: "Storage",
			Type: "sectorstorage.SealerConfig",

			Comment: ``,
		},
		{
			Name: "Fees",
			Type: "MinerFeeConfig",

			Comment: ``,
		},
		{
			Name: "Addresses",
			Type: "MinerAddressConfig",

			Comment: ``,
		},
		{
			Name: "DAGStore",
			Type: "DAGStoreConfig",

			Comment: ``,
		},
	},
	"Wallet": []DocField{
		{
			Name: "RemoteBackend",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "EnableLedger",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "DisableLocal",
			Type: "bool",

			Comment: ``,
		},
	},
}
