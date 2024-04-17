package adminusecase

import (
	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type AreaUsecase interface {
	RegisterArea(body *admindto.RegisterAreaRequest) (*dto.AreaKeyResponse, *dto.ErrorResponse)
	DeleteArea(query *admindto.DeleteAreaQuery) (*dto.AreaKeyResponse, *dto.ErrorResponse)
}

type areaUsecase struct {
	areaRepository repository.AreaRepository
}

func NewAreaUsecase(areaRepository repository.AreaRepository) AreaUsecase {
	return &areaUsecase{areaRepository: areaRepository}
}

func (u *areaUsecase) RegisterArea(body *admindto.RegisterAreaRequest) (*dto.AreaKeyResponse, *dto.ErrorResponse) {
	area := new(model.Area)
	area.AreaKey = body.AreaKey
	area.AreaName = body.AreaName

	area, err := u.areaRepository.CreateArea(area)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.AreaKeyResponse{
		AreaKey: area.AreaKey,
	}, nil
}

func (u *areaUsecase) DeleteArea(query *admindto.DeleteAreaQuery) (*dto.AreaKeyResponse, *dto.ErrorResponse) {
	area, err := u.areaRepository.FindAreaByKey(query.AreaKey)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	u.areaRepository.DeleteArea(area)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.AreaKeyResponse{
		AreaKey: area.AreaKey,
	}, nil
}
