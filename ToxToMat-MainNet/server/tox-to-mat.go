package main

import (
	"swap/services"
)

func main() {

	services.InitTestNet()
	//services.InitMainNet()
	// Approve
	//blockchain.ApproveTox("0xC0f98D355dc08f8187cc742A2fCb00e62De60E2F")
	// 部署，初始化
	services.InitSwap()

}
