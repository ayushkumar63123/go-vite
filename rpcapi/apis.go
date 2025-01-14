package rpcapi

import (
	"github.com/vitelabs/go-vite"
	"github.com/vitelabs/go-vite/rpc"
	"github.com/vitelabs/go-vite/rpcapi/api"
	"github.com/vitelabs/go-vite/rpcapi/api/filters"
)

func Init(dir, lvl string, testApi_prikey, testApi_tti string, netId uint, dexAvailable *bool) {
	api.InitLog(dir, lvl)
	api.InitTestAPIParams(testApi_prikey, testApi_tti)
	api.InitConfig(netId, dexAvailable)
}

func GetApi(vite *vite.Vite, apiModule string) rpc.API {
	switch apiModule {
	// private IPC
	case "health":
		return rpc.API{
			Namespace: "health",
			Version:   "1.0",
			Service:   api.NewHealthApi(vite),
			Public:    true,
		}
	case "wallet":
		return rpc.API{
			Namespace: "wallet",
			Version:   "1.0",
			Service:   api.NewWalletApi(vite),
			Public:    false,
		}
	case "private_onroad":
		return rpc.API{
			Namespace: "onroad",
			Version:   "1.0",
			Service:   api.NewPrivateOnroadApi(vite),
			Public:    false,
		}
		// public  WS HTTP IPC

	case "pow":
		return rpc.API{
			Namespace: "pow",
			Version:   "1.0",
			Service:   api.NewPow(vite),
			Public:    true,
		}
	case "debug":
		return rpc.API{
			Namespace: "debug",
			Version:   "1.0",
			Service:   api.NewDeprecated(),
			Public:    true,
		}
	case "consensusGroup":
		return rpc.API{
			Namespace: "debug",
			Version:   "1.0",
			Service:   api.NewDeprecated(),
			Public:    true,
		}
	case "ledger":
		return rpc.API{
			Namespace: "ledger",
			Version:   "1.0",
			Service:   api.NewLedgerApi(vite),
			Public:    true,
		}
	case "public_onroad":
		return rpc.API{
			Namespace: "onroad",
			Version:   "1.0",
			Service:   api.NewPublicOnroadApi(vite),
			Public:    true,
		}
	case "net":
		return rpc.API{
			Namespace: "net",
			Version:   "1.0",
			Service:   api.NewNetApi(vite),
			Public:    true,
		}
	case "contract":
		return rpc.API{
			Namespace: "contract",
			Version:   "1.0",
			Service:   api.NewContractApi(vite),
			Public:    true,
		}
	case "register":
		return rpc.API{
			Namespace: "register",
			Version:   "1.0",
			Service:   api.NewRegisterApi(vite),
			Public:    true,
		}
	case "vote":
		return rpc.API{
			Namespace: "vote",
			Version:   "1.0",
			Service:   api.NewVoteApi(vite),
			Public:    true,
		}
	case "mintage":
		return rpc.API{
			Namespace: "mintage",
			Version:   "1.0",
			Service:   api.NewMintageAPI(vite),
			Public:    true,
		}
	case "pledge":
		return rpc.API{
			Namespace: "pledge",
			Version:   "1.0",
			Service:   api.NewQuotaApi(vite),
			Public:    true,
		}
	case "dexfund":
		return rpc.API{
			Namespace: "dexfund",
			Version:   "1.0",
			Service:   api.NewDexFundApi(vite),
			Public:    true,
		}
	case "dextrade":
		return rpc.API{
			Namespace: "dextrade",
			Version:   "1.0",
			Service:   api.NewDexTradeApi(vite),
			Public:    true,
		}
	case "dex":
		return rpc.API{
			Namespace: "dex",
			Version:   "1.0",
			Service:   api.NewDexApi(vite),
			Public:    true,
		}
	case "private_dex":
		return rpc.API{
			Namespace: "dex",
			Version:   "1.0",
			Service:   api.NewDexPrivateApi(vite),
			Public:    false,
		}
	case "tx":
		return rpc.API{
			Namespace: "tx",
			Version:   "1.0",
			Service:   api.NewTxApi(vite),
			Public:    true,
		}
	case "dashboard":
		return rpc.API{
			Namespace: "dashboard",
			Version:   "1.0",
			Service:   api.NewDashboardApi(vite),
			Public:    true,
		}
	case "subscribe":
		return rpc.API{
			Namespace: "subscribe",
			Version:   "1.0",
			Service:   filters.NewSubscribeApi(vite),
			Public:    true,
		}
	case "sbpstats":
		return rpc.API{
			Namespace: "sbpstats",
			Version:   "1.0",
			Service:   api.NewStatsApi(vite),
			Public:    true,
		}
	case "util":
		return rpc.API{
			Namespace: "util",
			Version:   "1.0",
			Service:   api.NewUtilApi(vite),
			Public:    true,
		}
	case "data":
		return rpc.API{
			Namespace: "data",
			Version:   "1.0",
			Service:   api.NewDataApi(vite),
			Public:    true,
		}
	case "ledgerdebug":
		return rpc.API{
			Namespace: "ledgerdebug",
			Version:   "1.0",
			Service:   api.NewLedgerDebugApi(vite),
			Public:    false,
		}
	case "miner":
		return rpc.API{
			Namespace: "miner",
			Version:   "1.0",
			Service:   api.NewMinerApi(vite),
			Public:    false,
		}
	default:
		return rpc.API{Namespace: apiModule}
	}
}

func GetApis(vite *vite.Vite, apiModules ...string) map[string]rpc.API {
	var apis = make(map[string]rpc.API, len(apiModules))
	for _, m := range apiModules {
		apis[m] = GetApi(vite, m)
	}
	return apis
}
func MergeApis(first map[string]rpc.API, second map[string]rpc.API) []rpc.API {
	resultMap := make(map[string]rpc.API)

	for md, api := range first {
		resultMap[md] = api
	}

	for md1, api1 := range second {
		if _, ok := resultMap[md1]; ok {
			continue
		} else {
			resultMap[md1] = api1
		}
	}

	var result []rpc.API
	for _, r := range resultMap {
		result = append(result, r)
	}
	return result
}

func GetPublicApis(vite *vite.Vite) map[string]rpc.API {
	return GetApis(vite, "ledger", "net", "contract", "util", "health")
}
