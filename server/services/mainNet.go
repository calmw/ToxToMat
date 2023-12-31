package services

import "swap/blockchain"

func InitMainNet() {

	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId:     9001,
		RPC:         "https://rpc-8.matchscan.io/",
		SwapAddress: "0x4D02Ee03980e92fEa8D8874D6602D5c7e653B1c9",
		ToxAddress:  "0x96397347Ea2beE29713Bc48749eB277D6A36A407",
		PrivateKey:  "73992c90ce55a7b625c57fa9f377ab35bb71be1cae7294889ef311ec3517d0a4",
	}
}
