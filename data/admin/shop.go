/*
Package admindata 管理API関連構造体
*/
package admindata

/*
RegisterShopRequestBody | 店舗登録リクエストボディ
@type Request
*/
type RegisterShopRequestBody struct {
	ServiceID    uint   `json:"serviceId" validate:"required"`
	Area         string `json:"area" validate:"required"`
	SSID         string `json:"ssid"`
	ShopName     string `json:"shopName" validate:"required"`
	Description  string `json:"description"`
	Address      string `json:"address" validate:"required"`
	ShopType     string `json:"shopType" validate:"required"`
	OpeningHours string `json:"openingHours"`
	SeatsNum     int    `json:"seatsNum"`
	HasPower     bool   `json:"hasPower"`
}
