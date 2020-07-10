/*
Package data 汎用的な構造体
*/
package data

/*
MessageResponse | メッセージを返却するレスポンス
@type Response
*/
type MessageResponse struct {
	Message string `json:"message"`
}

/*
AreaKeyResponse | エリアキーのみのレスポンス
@type Response
*/
type AreaKeyResponse struct {
	AreaKey string `json:"areaKey"`
}

/*
ServiceIDResponse | Wi-FiサービスIDのみのレスポンス
@type Response
*/
type ServiceIDResponse struct {
	ServiceID uint `json:"serviceId"`
}

/*
ShopIDResponse | 店舗IDのみのレスポンス
@type Response
*/
type ShopIDResponse struct {
	ShopID uint `json:"shopId"`
}

/*
ReviewIDResponse | レビューIDのみのレスポンス
@type Response
*/
type ReviewIDResponse struct {
	ReviewID uint `json:"reviewId"`
}
