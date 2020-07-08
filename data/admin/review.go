/*
Package admindata 管理API関連構造体
*/
package admindata

import "time"

/*
ReviewListingResponseElement | レビュー一覧取得レスポンス要素
@type Response
*/
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

/*
ReviewListingResponse | レビュー一覧取得レスポンス
@type Response
*/
type ReviewListingResponse struct {
	ReviewList []ReviewListingResponseElement `json:"reviewList"`
	Total      int                            `json:"total"`
}
