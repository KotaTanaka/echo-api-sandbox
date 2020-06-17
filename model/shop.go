/*
Package model モデル
*/
package model

import "github.com/jinzhu/gorm"

/*
Shop | Wi-Fi提供店舗モデル
*/
type Shop struct {
	gorm.Model
	ServiceID    uint
	AreaKey      string `gorm:"size:20"`
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
