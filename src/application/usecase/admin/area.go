package adminusecase

import (
	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/jinzhu/gorm"
)

type AreaUsecase interface {
	RegisterArea(body *admindto.RegisterAreaRequest) (*dto.AreaKeyResponse, *dto.ErrorResponse)
	DeleteArea(query *admindto.DeleteAreaQuery) (*dto.AreaKeyResponse, *dto.ErrorResponse)
}

type areaUsecase struct {
	db *gorm.DB
}

func NewAreaUsecase(db *gorm.DB) AreaUsecase {
	return &areaUsecase{db: db}
}

func (u areaUsecase) RegisterArea(body *admindto.RegisterAreaRequest) (*dto.AreaKeyResponse, *dto.ErrorResponse) {
	area := new(model.Area)
	area.AreaKey = body.AreaKey
	area.AreaName = body.AreaName

	u.db.Create(&area)

	return &dto.AreaKeyResponse{
		AreaKey: area.AreaKey,
	}, nil
}

func (u areaUsecase) DeleteArea(query *admindto.DeleteAreaQuery) (*dto.AreaKeyResponse, *dto.ErrorResponse) {
	area := model.Area{}
	u.db.Where("area_key = ?", query.AreaKey).Find(&area)
	u.db.Delete(&area)

	return &dto.AreaKeyResponse{
		AreaKey: area.AreaKey,
	}, nil
}
