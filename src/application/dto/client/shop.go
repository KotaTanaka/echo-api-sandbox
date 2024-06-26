package clientdto

type ShopListingResponseElement struct {
	ShopID       uint     `json:"shopId"`
	WifiName     string   `json:"wifiName"`
	ServiceLink  string   `json:"serviceLink"`
	ShopName     string   `json:"shopName"`
	AreaKey      string   `json:"areaKey"`
	Description  string   `json:"description"`
	Address      string   `json:"address"`
	Access       string   `json:"access"`
	SSID         []string `json:"SSID"`
	ShopType     string   `json:"shopType"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	HasPower     bool     `json:"hasPower"`
	ReviewCount  int64    `json:"reviewCount"`
	Average      float32  `json:"average"`
}

type ShopListingResponse struct {
	ShopList []ShopListingResponseElement `json:"shopList"`
	Total    int                          `json:"total"`
}
