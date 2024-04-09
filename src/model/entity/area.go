package entity

import "github.com/jinzhu/gorm"

type Area struct {
	gorm.Model
	AreaKey  string `gorm:"size:20"`
	AreaName string `gorm:"size:255"`
	Shops    []Shop
}
