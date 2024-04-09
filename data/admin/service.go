package admindata

import "time"

type RegisterServiceRequestBody struct {
	WifiName string `json:"wifiName" validate:"required"`
	Link     string `json:"link" validate:"required"`
}

type ServiceListingResponseElement struct {
	ServiceID uint   `json:"serviceId"`
	WifiName  string `json:"wifiName"`
	Link      string `json:"link"`
	ShopCount int    `json:"shopCount"`
}

type ServiceListingResponse struct {
	ServiceList []ServiceListingResponseElement `json:"serviceList"`
	Total       int                             `json:"total"`
}

type UpdateServiceRequestBody struct {
	WifiName string `json:"wifiName"`
	Link     string `json:"link"`
}

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
