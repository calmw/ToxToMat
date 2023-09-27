package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"swap/db"
	"swap/log"
	"swap/models"
)

type CityNodeConfigs struct {
	ChainId     int64
	RPC         string
	SwapAddress string
	ToxAddress  string
	PrivateKey  string
}

var CityNodeConfig CityNodeConfigs

func Client(c CityNodeConfigs) *ethclient.Client {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		panic("dail failed")
	}
	return client
}

func GetAuth(cli *ethclient.Client) (error, *bind.TransactOpts) {
	privateKeyEcdsa, err := crypto.HexToECDSA(CityNodeConfig.PrivateKey)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(CityNodeConfig.ChainId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}

	//gasLimit := uint64(21000)
	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     nil,
		Signer:    auth.Signer, // Method to use for signing the transaction (mandatory)
		Value:     big.NewInt(0),
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false, // Do all transact steps but do not send the transaction
	}
}

func GetStartBlock() uint64 {
	var blockStore = models.BlockStore{}
	db.Mysql.Table("block_store").Where("id=1").First(&blockStore)
	return blockStore.StartBlock
}

func SetSTartBlock(startBlock int64) {
	db.Mysql.Table("block_store").Where("id=1").Update("start_block", startBlock)
}
