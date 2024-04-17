package clientusecase

import (
	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type AreaUsecase interface {
	GetAreaMaster() (*clientdto.AreaMasterResponse, *dto.ErrorResponse)
}

type areaUsecase struct {
	areaRepository repository.AreaRepository
}

func NewAreaUsecase(areaRepository repository.AreaRepository) AreaUsecase {
	return &areaUsecase{areaRepository: areaRepository}
}

func (u *areaUsecase) GetAreaMaster() (*clientdto.AreaMasterResponse, *dto.ErrorResponse) {
	areas, err := u.areaRepository.ListAreas()
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	var res *clientdto.AreaMasterResponse
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
