/*
Package data | 店舗関連の構造体
*/
package data

/*
ShopIDResponse | 店舗IDのみのレスポンス
@type Response
*/
type ShopIDResponse struct {
	ShopID uint `json:"shopId"`
}

/*
RegisterShopRequestBody | 店舗登録リクエストボディ
@type Request
*/
type RegisterShopRequestBody struct {
	ServiceID    uint   `json:"serviceId" validate:"required"`
	SSID         string `json:"ssid"`
	ShopName     string `json:"shopName" validate:"required"`
	Description  string `json:"description"`
	Address      string `json:"address" validate:"required"`
	ShopType     string `json:"shopType" validate:"required"`
	OpeningHours string `json:"openingHours"`
	SeatsNum     int    `json:"seatsNum"`
	HasPower     bool   `json:"hasPower"`
}

/*
ShopListingResponseElement | 店舗一覧取得レスポンス要素
@type Response
*/
type ShopListingResponseElement struct {
	ShopID       uint     `json:"shopId"`
	ShopName     string   `json:"shopName"`
	WifiName     string   `json:"wifiName"`
	ServiceLink  string   `json:"serviceLink"`
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
ShopListingResponse | 店舗一覧取得レスポンス
@type Response
*/
type ShopListingResponse struct {
	ShopList []ShopListingResponseElement `json:"shopList"`
	Total    int                          `json:"total"`
}
