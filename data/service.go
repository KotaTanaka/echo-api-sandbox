/*
Package data --- Wi-Fiサービス関連の構造体
*/
package data

import "github.com/jinzhu/gorm"

/*
Service --- Model Wi-Fiサービステーブル
*/
type Service struct {
	gorm.Model
	WifiName string `gorm:"size:255"`
	Link     string `gorm:"size:255"`
	Shops    []Shop
}

/*
ServiceIDResponse --- Wi-FiサービスIDのみのレスポンス
*/
type ServiceIDResponse struct {
	ServiceID uint `json:"serviceId"`
}

/*
RegisterServiceRequestBody --- Wi-Fiサービス登録リクエストボディ
*/
type RegisterServiceRequestBody struct {
	WifiName string `json:"wifiName" validate:"required"`
	Link     string `json:"link" validate:"required"`
}

/*
ServiceListingResponseElement ---  Wi-Fiサービス一覧取得レスポンス要素
*/
type ServiceListingResponseElement struct {
	ServiceID uint   `json:"serviceId"`
	WifiName  string `json:"wifiName"`
	Link      string `json:"Link"`
	ShopCount int    `json:"shopCount"`
}

/*
ServiceListingResponse ---  Wi-Fiサービス一覧取得レスポンス
*/
type ServiceListingResponse struct {
	ServiceList []ServiceListingResponseElement `json:"serviceList"`
	Total       int                             `json:"total"`
}
