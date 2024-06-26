package admindto

import "time"

type RegisterShopRequest struct {
	ServiceID    uint     `json:"serviceId" validate:"required"`
	ShopName     string   `json:"shopName" validate:"required"`
	AreaID       uint     `json:"areaId" validate:"required"`
	Description  string   `json:"description"`
	Address      string   `json:"address" validate:"required"`
	Access       string   `json:"access"`
	SSID         []string `json:"ssid" validate:"required"`
	ShopType     string   `json:"shopType"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	HasPower     bool     `json:"hasPower"`
}

type UpdateShopRequest struct {
	ShopName     string   `json:"shopName"`
	AreaID       uint     `json:"areaId"`
	Description  string   `json:"description"`
	Address      string   `json:"address"`
	Access       string   `json:"access"`
	SSID         []string `json:"ssid"`
	ShopType     string   `json:"shopType"`
	OpeningHours string   `json:"openingHours"`
	SeatsNum     int      `json:"seatsNum"`
	HasPower     bool     `json:"hasPower"`
}

type ShopListingResponseElement struct {
	ShopID       uint     `json:"shopId"`
	ServiceID    uint     `json:"serviceId"`
	WifiName     string   `json:"wifiName"`
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

type ShopDetailResponse struct {
	ShopID       uint                                  `json:"shopId"`
	ServiceID    uint                                  `json:"serviceId"`
	WifiName     string                                `json:"wifiName"`
	ShopName     string                                `json:"shopName"`
	AreaKey      string                                `json:"areaKey"`
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
	ReviewCount  int64                                 `json:"reviewCount"`
	ReviewList   []ShopDetailResponseReviewListElement `json:"reviewList"`
	Average      float32                               `json:"average"`
}

type ShopDetailResponseReviewListElement struct {
	ReviewID   uint       `json:"reviewId"`
	Comment    string     `json:"comment"`
	Evaluation int        `json:"evaluation"`
	Status     bool       `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}
