package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	ListServices() ([]*model.Service, error)
	FindServiceByID(serviceID int) (*model.Service, error)
	CreateService(service *model.Service) (*model.Service, error)
	UpdateService(service *model.Service) (*model.Service, error)
	DeleteService(service *model.Service) error
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) ListServices() ([]*model.Service, error) {
	services := []*model.Service{}
	if err := r.db.
		Preload("Shops").
		Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (r *serviceRepository) FindServiceByID(serviceID int) (*model.Service, error) {
	var service model.Service
	if err := r.db.
		Preload("Shops.Area").
		First(&service, serviceID).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (r *serviceRepository) CreateService(service *model.Service) (*model.Service, error) {
	if err := r.db.Create(&service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (r *serviceRepository) UpdateService(service *model.Service) (*model.Service, error) {
	if err := r.db.Save(&service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (r *serviceRepository) DeleteService(service *model.Service) error {
	return r.db.Delete(&service).Error
}
