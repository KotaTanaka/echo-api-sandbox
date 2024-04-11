package admindto

type RegisterAreaRequest struct {
	AreaKey  string `json:"areaKey" validate:"required"`
	AreaName string `json:"areaName" validate:"required"`
}

type DeleteAreaQuery struct {
	AreaKey string
}
