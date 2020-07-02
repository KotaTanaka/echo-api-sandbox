/*
Package clientdata クライアントAPI関連構造体
*/
package clientdata

/*
AreaMasterResponseElement | エリアマスタ取得レスポンス要素
@type Response
*/
type AreaMasterResponseElement struct {
	AreaKey   string `json:"areaKey"`
	AreaName  string `json:"areaName"`
	ShopCount int    `json:"shopCount"`
}

/*
AreaMasterResponse | エリアマスタ取得レスポンス
@type Response
*/
type AreaMasterResponse struct {
	AreaList []AreaMasterResponseElement `json:"areaList"`
}
