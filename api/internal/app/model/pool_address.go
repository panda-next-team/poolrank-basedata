package model

import "time"

type PoolAddress struct {
	Id          int32 `xorm:"pk autoincr"`
	PoolId      int32
	CoinId      int32
	Address     string
	Type        int32
	CreatedAtTs int32
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAtTs int32
	UpdatedAt   time.Time `xorm:"updated"`
}
