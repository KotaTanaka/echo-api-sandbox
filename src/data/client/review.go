package clientdata

import "time"

type ReviewListingResponseElement struct {
	ReviewID   uint      `json:"reviewId"`
	Comment    string    `json:"comment"`
	Evaluation int       `json:"evaluation"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ReviewListingResponse struct {
	ShopID     uint                           `json:"shopId"`
	ShopName   string                         `json:"shopName"`
	ServiceID  uint                           `json:"serviceId"`
	WifiName   string                         `json:"wifiName"`
	Average    float32                        `json:"average"`
	ReviewList []ReviewListingResponseElement `json:"reviewList"`
	Total      int                            `json:"total"`
}

type CreateReviewRequestBody struct {
	ShopID     uint   `json:"shopId" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
	Evaluation int    `json:"evaluation" validate:"required"`
}
