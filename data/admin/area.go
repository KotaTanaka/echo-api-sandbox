/*
Package admindata 管理API関連構造体
*/
package admindata

/*
RegisterAreaRequestBody | エリア登録リクエストボディ
@type Request
*/
type RegisterAreaRequestBody struct {
	AreaKey  string `json:"areaKey" validate:"required"`
	AreaName string `json:"areaName" validate:"required"`
}
