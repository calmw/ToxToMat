package services

import (
	"swap/blockchain"
)

func InitSwap() {
	// 管理员设置TOX合约地址
	blockchain.AdminSetToxAddress()
	// 管理员设置每次兑换，最多使用的TOX数量
	blockchain.AdminSetMaxToxPerTime(100)
	// 管理员设置每次兑换，最多使用的TOX数量
	blockchain.AdminSetMinToxPerTime(10)
	// 管理员设置兑换汇率
	blockchain.AdminSetRate(100)
	// 管理员设置用户，最最多可使用的TOX数量
	blockchain.AdminSetSwapTotalMax(100)

}
