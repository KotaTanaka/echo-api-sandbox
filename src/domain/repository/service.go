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
	r.db.Preload("Shops").Find(&services)

	return services, nil
}

func (r *serviceRepository) FindServiceByID(serviceID int) (*model.Service, error) {
	var service model.Service
	r.db.Preload("Shops.Area").First(&service, serviceID)

	return &service, nil
}

func (r *serviceRepository) CreateService(service *model.Service) (*model.Service, error) {
	r.db.Create(&service)

	return service, nil
}

func (r *serviceRepository) UpdateService(service *model.Service) (*model.Service, error) {
	r.db.Save(&service)

	return service, nil
}

func (r *serviceRepository) DeleteService(service *model.Service) error {
	r.db.Delete(&service)

	return nil
}
