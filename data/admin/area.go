package admindata

type RegisterAreaRequestBody struct {
	AreaKey  string `json:"areaKey" validate:"required"`
	AreaName string `json:"areaName" validate:"required"`
}
