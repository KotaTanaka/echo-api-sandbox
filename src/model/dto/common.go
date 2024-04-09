package dto

type MessageResponse struct {
	Message string `json:"message"`
}

type AreaKeyResponse struct {
	AreaKey string `json:"areaKey"`
}

type ServiceIDResponse struct {
	ServiceID uint `json:"serviceId"`
}

type ShopIDResponse struct {
	ShopID uint `json:"shopId"`
}

type ReviewIDResponse struct {
	ReviewID uint `json:"reviewId"`
}
