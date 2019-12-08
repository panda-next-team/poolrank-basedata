package model

import "time"

type Algorithm struct {
	Id          int32 `xorm:"pk autoincr"`
	Name        string
	CreatedAtTs int64
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAtTs int64
	UpdatedAt   time.Time `xorm:"updated"`
}
