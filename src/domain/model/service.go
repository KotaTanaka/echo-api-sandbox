package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	WifiName string `gorm:"size:255"`
	Link     string `gorm:"size:255"`
	Shops    []Shop
}
