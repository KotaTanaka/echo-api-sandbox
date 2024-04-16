package clientusecase

import (
	"github.com/jinzhu/gorm"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type AreaUsecase interface {
	GetAreaMaster() (*clientdto.AreaMasterResponse, *dto.ErrorResponse)
}

type areaUsecase struct {
	db *gorm.DB
}

func NewAreaUsecase(db *gorm.DB) AreaUsecase {
	return &areaUsecase{db: db}
}

func (u *areaUsecase) GetAreaMaster() (*clientdto.AreaMasterResponse, *dto.ErrorResponse) {
	areas := []model.Area{}
	u.db.Find(&areas)

	res := &clientdto.AreaMasterResponse{}
	res.AreaList = []clientdto.AreaMasterResponseElement{}

	for _, area := range areas {
		res.AreaList = append(
			res.AreaList,
			clientdto.AreaMasterResponseElement{
				AreaKey:   area.AreaKey,
				AreaName:  area.AreaName,
				ShopCount: len(area.Shops),
			},
		)
	}

	return res, nil
}
