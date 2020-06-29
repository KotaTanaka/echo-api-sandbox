/*
Package clientdata クライアントAPI関連構造体
*/
package clientdata

/*
ShopListingResponseElement | 店舗一覧取得レスポンス要素
@type Response
*/
type ShopListingResponseElement struct {
	ShopID       uint     `json:"shopId"`
	WifiName     string   `json:"wifiName"`
	ServiceLink  string   `json:"serviceLink"`
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
ShopListingResponse | 店舗一覧取得レスポンス
@type Response
*/
type ShopListingResponse struct {
	ShopList []ShopListingResponseElement `json:"shopList"`
	Total    int                          `json:"total"`
}
