package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	intoCityNode "swap/binding"
	"swap/log"
)

func ApproveTox(contractAddress string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	tox, err := intoCityNode.NewBgtToken(common.HexToAddress(CityNodeConfig.ToxAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amount := E18.Mul(E18, big.NewInt(100000))
	tx, err := tox.Approve(auth, common.HexToAddress(contractAddress), amount)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(tx, err)
}
