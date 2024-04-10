package clientdto

type AreaMasterResponseElement struct {
	AreaKey   string `json:"areaKey"`
	AreaName  string `json:"areaName"`
	ShopCount int    `json:"shopCount"`
}

type AreaMasterResponse struct {
	AreaList []AreaMasterResponseElement `json:"areaList"`
}
