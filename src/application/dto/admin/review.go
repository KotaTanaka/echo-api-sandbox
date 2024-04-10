package admindto

import "time"

type UpdateReviewStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type ReviewListingResponseElement struct {
	ReviewID   uint       `json:"reviewId"`
	ShopID     uint       `json:"shopId"`
	ShopName   string     `json:"shopName"`
	ServiceID  uint       `json:"serviceId"`
	WifiName   string     `json:"wifiName"`
	Comment    string     `json:"comment"`
	Evaluation int        `json:"evaluation"`
	Status     bool       `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

type ReviewListingResponse struct {
	ReviewList []ReviewListingResponseElement `json:"reviewList"`
	Total      int                            `json:"total"`
}
