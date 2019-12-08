package model

import "time"

type Pool struct {
	Id          int32 `xorm:"pk autoincr"`
	Name        string
	WebsiteUrl  string
	Icon        string
	Status      int32
	ListOrder   int32
	CreatedAtTs int32
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAtTs int32
	UpdatedAt   time.Time `xorm:"updated"`
}
