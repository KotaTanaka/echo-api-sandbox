/*
Package admindata 管理API関連構造体
*/
package admindata

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
