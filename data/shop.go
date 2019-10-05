/*
Package data - Shop関連の構造体
*/
package data

import (
	"github.com/jinzhu/gorm"
)

/*
Shop Model
*/
type Shop struct {
	gorm.Model
	ServiceID    int
	SSID         string `gorm:"size:255"`
	Name         string `gorm:"size:255"`
	Descripttion string `gorm:"size:255"`
	Address      string `gorm:"size:255"`
	ShopType     string `gorm:"size:255"`
	OpeningHours string `gorm:"size:255"`
	SeatsNum     int
	hasPower     bool
}

/*
ShopListingResponseElement - 店舗一覧取得レスポンス要素
*/
type ShopListingResponseElement struct {
	ShopID       string   `json:"shopId"`
	ShopName     string   `json:"shopName"`
	WifiName     string   `json:"wifiName"`
	Ssid         []string `json:"ssid"`
	Address      string   `json:"address"`
	Acceess      string   `json:"access"`
	Description  string   `json:"description"`
	ShopType     string   `json:"shoptype"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	Power        bool     `json:"power"`
	ReviewCount  int      `json:"reviewCount"`
	Average      float32  `json:"average"`
}

/*
ShopListingResponse - 店舗一覧取得レスポンス
*/
type ShopListingResponse struct {
	ShopList []ShopListingResponseElement `json:"shopList"`
	Total    int                          `json:"total"`
}
