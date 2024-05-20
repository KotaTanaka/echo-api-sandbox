package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ServiceID    uint
	Service      Service
	AreaID       uint
	Area         Area
	ShopName     string `gorm:"size:255"`
	Description  string `gorm:"size:255"`
	Address      string `gorm:"size:255"`
	Access       string `gorm:"size:255"`
	SSID         string `gorm:"size:255"`
	ShopType     string `gorm:"size:255"`
	OpeningHours string `gorm:"size:255"`
	SeatsNum     int
	HasPower     bool
	Reviews      []Review
}
