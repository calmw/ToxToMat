package services

import "swap/blockchain"

func InitTestNet() {

	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId:     9001,
		RPC:         "https://testnet-rpc.d2ao.com/",
		SwapAddress: "0xC0f98D355dc08f8187cc742A2fCb00e62De60E2F",
		ToxAddress:  "0x3eE243ff68074502b1D9D65443fa97b99f634570",
		PrivateKey:  "73992c90ce55a7b625c57fa9f377ab35bb71be1cae7294889ef311ec3517d0a4",
	}
}
