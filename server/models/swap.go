package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Swap struct {
	Id          int             `gorm:"column:id;primaryKey"`
	User        string          `json:"user" gorm:"column:user"`
	TxHash      string          `json:"tx_hash" gorm:"column:tx_hash"`
	Tox         decimal.Decimal `json:"tox" gorm:"column:tox"`
	Mat         decimal.Decimal `json:"mat" gorm:"column:mat"`
	BlockHeight int64           `json:"block_height" gorm:"column:block_height"`
	Ctime       time.Time       `json:"ctime" gorm:"column:ctime"`
}
