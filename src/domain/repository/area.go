package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/jinzhu/gorm"
)

type AreaRepository interface {
	CreateArea(area *model.Area) (*model.Area, error)
	ListAreas() ([]*model.Area, error)
	FindAreaByKey(areaKey string) (*model.Area, error)
	DeleteArea(area *model.Area) error
}

type areaRepository struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) AreaRepository {
	return &areaRepository{db: db}
}

func (r *areaRepository) CreateArea(area *model.Area) (*model.Area, error) {
	r.db.Create(&area)

	return area, nil
}

func (r *areaRepository) ListAreas() ([]*model.Area, error) {
	areas := []*model.Area{}
	r.db.Find(&areas)

	return areas, nil
}

func (r *areaRepository) FindAreaByKey(areaKey string) (*model.Area, error) {
	var area *model.Area
	r.db.Where("area_key = ?", areaKey).Find(&area)

	return area, nil
}

func (r *areaRepository) DeleteArea(area *model.Area) error {
	r.db.Delete(&area)

	return nil
}
