/*
Package admindata 管理API関連構造体
*/
package admindata

import "time"

/*
RegisterShopRequestBody | 店舗登録リクエストボディ
@type Request
*/
type RegisterShopRequestBody struct {
	ServiceID    uint     `json:"serviceId" validate:"required"`
	ShopName     string   `json:"shopName" validate:"required"`
	Area         string   `json:"area" validate:"required"`
	Description  string   `json:"description"`
	Address      string   `json:"address" validate:"required"`
	Access       string   `json:"access"`
	SSID         []string `json:"ssid" validate:"required"`
	ShopType     string   `json:"shopType"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	HasPower     bool     `json:"hasPower"`
}

/*
UpdateShopRequestBody | 店舗編集リクエストボディ
@type Request
*/
type UpdateShopRequestBody struct {
	ShopName     string   `json:"shopName"`
	Area         string   `json:"area"`
	Description  string   `json:"description"`
	Address      string   `json:"address"`
	Access       string   `json:"access"`
	SSID         []string `json:"ssid"`
	ShopType     string   `json:"shopType"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	HasPower     bool     `json:"hasPower"`
}

/*
ShopListingResponseElement | 店舗一覧取得・検索レスポンス要素
@type Response
*/
type ShopListingResponseElement struct {
	ShopID       uint     `json:"shopId"`
	ServiceID    uint     `json:"serviceId"`
	WifiName     string   `json:"wifiName"`
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

/*
ShopListingResponse | 店舗一覧取得・検索レスポンス
@type Response
*/
type ShopListingResponse struct {
	ShopList []ShopListingResponseElement `json:"shopList"`
	Total    int                          `json:"total"`
}

/*
ShopDetailResponse | 店舗詳細取得レスポンス
@type Response
*/
type ShopDetailResponse struct {
	ShopID       uint                                  `json:"shopId"`
	ServiceID    uint                                  `json:"serviceId"`
	WifiName     string                                `json:"wifiName"`
	ShopName     string                                `json:"shopName"`
	Area         string                                `json:"area"`
	Description  string                                `json:"description"`
	Address      string                                `json:"address"`
	Access       string                                `json:"access"`
	SSID         []string                              `json:"SSID"`
	ShopType     string                                `json:"shopType"`
	OpeningHours string                                `json:"openingHours"`
	SeatsNum     int                                   `json:"seatsNum"`
	HasPower     bool                                  `json:"hasPower"`
	CreatedAt    time.Time                             `json:"createdAt"`
	UpdatedAt    time.Time                             `json:"updatedAt"`
	DeletedAt    *time.Time                            `json:"deletedAt"`
	ReviewCount  int                                   `json:"reviewCount"`
	ReviewList   []ShopDetailResponseReviewListElement `json:"reviewList"`
	Average      float32                               `json:"average"`
}

/*
ShopDetailResponseReviewListElement | 店舗詳細取得レスポンスレビューリスト要素
@type Response
*/
type ShopDetailResponseReviewListElement struct {
	ReviewID   uint       `json:"reviewId"`
	Comment    string     `json:"comment"`
	Evaluation int        `json:"evaluation"`
	Status     bool       `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}
