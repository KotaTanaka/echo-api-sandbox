/*
Package data | サービス関連の構造体
*/
package data

/*
ServiceIDResponse | Wi-FiサービスIDのみのレスポンス
@type Response
*/
type ServiceIDResponse struct {
	ServiceID uint `json:"serviceId"`
}

/*
RegisterServiceRequestBody | Wi-Fiサービス登録リクエストボディ
@type Request
*/
type RegisterServiceRequestBody struct {
	WifiName string `json:"wifiName" validate:"required"`
	Link     string `json:"link" validate:"required"`
}

/*
ServiceListingResponseElement | Wi-Fiサービス一覧取得レスポンス要素
@type Response
*/
type ServiceListingResponseElement struct {
	ServiceID uint   `json:"serviceId"`
	WifiName  string `json:"wifiName"`
	Link      string `json:"link"`
	ShopCount int    `json:"shopCount"`
}

/*
ServiceListingResponse | Wi-Fiサービス一覧取得レスポンス
@type Response
*/
type ServiceListingResponse struct {
	ServiceList []ServiceListingResponseElement `json:"serviceList"`
	Total       int                             `json:"total"`
}
