package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/big"
	"strings"
	intoSwap "swap/binding"
	"swap/blockchain/event"
	"swap/db"
	"swap/log"
	"swap/models"
	"time"
)

// AdminSetToxAddress 管理员设置TOX合约地址
func AdminSetToxAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoSwap.NewSwap(common.HexToAddress(CityNodeConfig.SwapAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, err = city.AdminSetToxAddress(auth, common.HexToAddress(CityNodeConfig.ToxAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// AdminSetMaxToxPerTime 管理员设置每次兑换，最多使用的TOX数量
func AdminSetMaxToxPerTime(amount int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoSwap.NewSwap(common.HexToAddress(CityNodeConfig.SwapAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amountBigInt := E18.Mul(E18, big.NewInt(amount))
	_, err = city.AdminSetMaxToxPerTime(auth, amountBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// AdminSetMinToxPerTime 管理员设置每次兑换，最多使用的TOX数量
func AdminSetMinToxPerTime(amount int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoSwap.NewSwap(common.HexToAddress(CityNodeConfig.SwapAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amountBigInt := E18.Mul(E18, big.NewInt(amount))
	_, err = city.AdminSetMinToxPerTime(auth, amountBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// AdminSetRate 管理员设置兑换汇率
func AdminSetRate(rate int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoSwap.NewSwap(common.HexToAddress(CityNodeConfig.SwapAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, err = city.AdminSetRate(auth, big.NewInt(rate))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// AdminSetSwapTotalMax 管理员设置用户，最最多可使用的TOX数量
func AdminSetSwapTotalMax(swapTotalMax int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoSwap.NewSwap(common.HexToAddress(CityNodeConfig.SwapAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	swapTotalMaxBigInt := E18.Mul(E18, big.NewInt(swapTotalMax))
	_, err = city.AdminSetSwapTotalMax(auth, swapTotalMaxBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

func GetSwapRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	InsertDailyRewardLock.Lock()
	defer InsertDailyRewardLock.Unlock()
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.SwapAddress),
		event.SwapRecordEvent.EventSignature,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	logs, err := Cli.FilterLogs(ctx, query)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cancel()
	abi, _ := intoSwap.SwapMetaData.GetAbi()

	for _, logE := range logs {
		logData, err := abi.Unpack(event.SwapRecordEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		var timestamp int64
		var txHash string
		userAddress := strings.ToLower(logData[0].(common.Address).String())
		toxNum := decimal.NewFromBigInt(logData[1].(*big.Int), 0)
		matNum := decimal.NewFromBigInt(logData[2].(*big.Int), 0)
		ctime := logData[3].(*big.Int)
		header, err := Cli.HeaderByNumber(context.Background(), big.NewInt(int64(logE.BlockNumber)))
		if err == nil {
			timestamp = int64(header.Time)
		}
		if header != nil {
			txHash = header.TxHash.String()
		}
		fmt.Println(header)
		fmt.Println(txHash)
		fmt.Println(int64(logE.BlockNumber))
		fmt.Println(timestamp)
		err = InsertSwapRecord(userAddress, txHash, toxNum, matNum, int64(logE.BlockNumber), ctime.Int64())
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertSwapRecord(userAddress, txHash string, toxNum, matNum decimal.Decimal, blockHeight, timestamp int64) error {

	var swap models.Swap
	whereCondition := fmt.Sprintf("user='%s' and block_height=%d and tx_hash='%s'", strings.ToLower(userAddress), blockHeight, strings.ToLower(txHash))
	err := db.Mysql.Table("reward").Where(whereCondition).First(&swap).Error
	if err == gorm.ErrRecordNotFound {
		// 插入数据
		db.Mysql.Model(&models.Swap{}).Create(&models.Swap{
			User:        strings.ToLower(userAddress),
			TxHash:      strings.ToLower(txHash),
			Tox:         toxNum,
			Mat:         matNum,
			BlockHeight: blockHeight,
			Ctime:       time.Unix(timestamp, 0),
		})
	}
	return nil
}
