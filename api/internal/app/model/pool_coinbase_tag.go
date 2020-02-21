package model

import "time"

type PoolCoinbaseTag struct {
	Id          int32 `xorm:"pk autoincr"`
	PoolId      int32
	Tag         []byte
	CreatedAtTs int32
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAtTs int32
	UpdatedAt   time.Time `xorm:"updated"`
}
