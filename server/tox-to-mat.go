package main

import (
	"swap/blockchain"
	"swap/services"
)

func main() {

	services.InitTestNet()
	// Approve
	blockchain.ApproveTox("0xC0f98D355dc08f8187cc742A2fCb00e62De60E2F")

}
