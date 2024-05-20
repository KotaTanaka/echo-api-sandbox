package model

import "gorm.io/gorm"

type Area struct {
	gorm.Model
	AreaKey  string `gorm:"size:20"`
	AreaName string `gorm:"size:255"`
	Shops    []Shop
}
