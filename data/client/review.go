/*
Package clientdata クライアントAPI関連構造体
*/
package clientdata

import "time"

/*
ReviewListingResponseElement | レビュー一覧取得レスポンス要素
@type Response
*/
type ReviewListingResponseElement struct {
	ReviewID   uint      `json:"reviewId"`
	Comment    string    `json:"comment"`
	Evaluation int       `json:"evaluation"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

/*
ReviewListingResponse | レビュー一覧取得レスポンス
@type Response
*/
type ReviewListingResponse struct {
	ShopID     uint                           `json:"shopId"`
	ShopName   string                         `json:"shopName"`
	ServiceID  uint                           `json:"serviceId"`
	WifiName   string                         `json:"wifiName"`
	Average    float32                        `json:"average"`
	ReviewList []ReviewListingResponseElement `json:"reviewList"`
	Total      int                            `json:"total"`
}

/*
CreateReviewRequestBody | レビュー投稿リクエストボディ
@type Request
*/
type CreateReviewRequestBody struct {
	ShopID     uint   `json:"shopId" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
	Evaluation int    `json:"evaluation" validate:"required"`
}
