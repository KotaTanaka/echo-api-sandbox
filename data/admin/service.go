/*
Package admindata 管理API関連構造体
*/
package admindata

import "time"

/*
RegisterServiceRequestBody | Wi-Fiサービス登録リクエストボディ
@type Request
*/
type RegisterServiceRequestBody struct {
	WifiName string `json:"wifiName" validate:"required"`
	Link     string `json:"link" validate:"required"`
}

/*
ServiceListingResponseElement | Wi-Fiサービス一覧取得・検索レスポンス要素
@type Response
*/
type ServiceListingResponseElement struct {
	ServiceID uint   `json:"serviceId"`
	WifiName  string `json:"wifiName"`
	Link      string `json:"link"`
	ShopCount int    `json:"shopCount"`
}

/*
ServiceListingResponse | Wi-Fiサービス一覧取得レスポンス
@type Response
*/
type ServiceListingResponse struct {
	ServiceList []ServiceListingResponseElement `json:"serviceList"`
	Total       int                             `json:"total"`
}

/*
UpdateServiceRequestBody | Wi-Fiサービス編集リクエストボディ
@type Request
*/
type UpdateServiceRequestBody struct {
	WifiName string `json:"wifiName"`
	Link     string `json:"link"`
}

/*
ServiceDetailResponse | Wi-Fiサービス詳細取得レスポンス
@type Response
*/
type ServiceDetailResponse struct {
	ServiceID uint                                   `json:"serviceId"`
	WifiName  string                                 `json:"wifiName"`
	Link      string                                 `json:"link"`
	CreatedAt time.Time                              `json:"createdAt"`
	UpdatedAt time.Time                              `json:"updatedAt"`
	DeletedAt *time.Time                             `json:"deletedAt"`
	ShopCount int                                    `json:"shopCount"`
	ShopList  []ServiceDetailResponseShopListElement `json:"shopList"`
}

/*
ServiceDetailResponseShopListElement |  Wi-Fiサービス詳細取得レスポンス店舗リスト要素
@type Response
*/
type ServiceDetailResponseShopListElement struct {
	ShopID       uint     `json:"shopId"`
	ShopName     string   `json:"shopName"`
	Area         string   `json:"area"`
	Description  string   `json:"description"`
	Address      string   `json:"address"`
	Access       string   `json:"access"`
	SSID         []string `json:"SSID"`
	ShopType     string   `json:"shopType"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	HasPower     bool     `json:"hasPower"`
	ReviewCount  int      `json:"reviewCount"`
	Average      float32  `json:"average"`
}
