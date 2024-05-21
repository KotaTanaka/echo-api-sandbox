package adminusecase

import (
	"fmt"

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
	area := &model.Area{
		AreaKey:  body.AreaKey,
		AreaName: body.AreaName,
	}

	area, err := u.areaRepository.CreateArea(area)
	if err != nil {
		return nil, dto.HandleDBError(err, "Area")
	}

	return &dto.AreaKeyResponse{
		AreaKey: area.AreaKey,
	}, nil
}

func (u *areaUsecase) DeleteArea(query *admindto.DeleteAreaQuery) (*dto.AreaKeyResponse, *dto.ErrorResponse) {
	area, err := u.areaRepository.FindAreaByKey(query.AreaKey)
	if err != nil {
		return nil, dto.HandleDBError(err, fmt.Sprintf("Area(key:%s)", query.AreaKey))
	}

	if err := u.areaRepository.DeleteArea(area); err != nil {
		return nil, dto.HandleDBError(err, fmt.Sprintf("Area(key:%s)", area.AreaKey))
	}

	return &dto.AreaKeyResponse{
		AreaKey: area.AreaKey,
	}, nil
}
