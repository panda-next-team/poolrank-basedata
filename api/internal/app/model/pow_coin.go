package model

import "time"

type PowCoin struct {
	Id          int32  `xorm:"pk autoincr"`
	Name        string `xorm:"unique"`
	EnName      string `xorm:"unique"`
	EnTag       string `xorm:"unique"`
	AlgorithmId int32
	MaxSupply   float64
	ReleaseDate time.Time
	BlockTime   int32
	Icon        string
	GithubUrl   string
	WebsiteUrl  string
	Intro       string
	Status      int32
	ListOrder   int32
	CreatedAtTs int32
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAtTs int32
	UpdatedAt   time.Time `xorm:"updated"`
}
