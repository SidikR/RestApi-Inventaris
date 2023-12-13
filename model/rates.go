package model

type Rates struct {
	ID      uint `gorm:"primary_key"`
	Service uint `gorm:"default:100000"`
}
