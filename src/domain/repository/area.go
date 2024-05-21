package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"gorm.io/gorm"
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
	if err := r.db.Create(&area).Error; err != nil {
		return nil, err
	}

	return area, nil
}

func (r *areaRepository) ListAreas() ([]*model.Area, error) {
	areas := []*model.Area{}
	if err := r.db.
		Preload("Shops").
		Find(&areas).Error; err != nil {
		return nil, err
	}

	return areas, nil
}

func (r *areaRepository) FindAreaByKey(areaKey string) (*model.Area, error) {
	var area model.Area
	if err := r.db.
		Where("area_key = ?", areaKey).
		First(&area).Error; err != nil {
		return nil, err
	}

	return &area, nil
}

func (r *areaRepository) DeleteArea(area *model.Area) error {
	return r.db.Delete(&area).Error
}
