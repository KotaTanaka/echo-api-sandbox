package adminusecase

import (
	"strings"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/jinzhu/gorm"
)

type ServiceUsecase interface {
	GetServiceList() (*admindto.ServiceListingResponse, *dto.ErrorResponse)
	GetServiceDetail(serviceID int) (*admindto.ServiceDetailResponse, *dto.ErrorResponse)
	RegisterService(body *admindto.RegisterServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse)
	UpdateService(serviceID int, body *admindto.UpdateServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse)
	DeleteService(serviceID int) (*dto.ServiceIDResponse, *dto.ErrorResponse)
}

type serviceUsecase struct {
	db *gorm.DB
}

func NewServiceUsecase(db *gorm.DB) ServiceUsecase {
	return &serviceUsecase{db: db}
}

func (u *serviceUsecase) GetServiceList() (*admindto.ServiceListingResponse, *dto.ErrorResponse) {
	services := []model.Service{}
	u.db.Find(&services)

	res := &admindto.ServiceListingResponse{}
	res.Total = len(services)
	res.ServiceList = []admindto.ServiceListingResponseElement{}

	for _, service := range services {
		shopCount := 0
		u.db.Model(&model.Shop{}).Where("service_id = ?", service.ID).Count(&shopCount)

		res.ServiceList = append(
			res.ServiceList, admindto.ServiceListingResponseElement{
				ServiceID: service.ID,
				WifiName:  service.WifiName,
				Link:      service.Link,
				ShopCount: shopCount,
			},
		)
	}

	return res, nil
}

func (u *serviceUsecase) GetServiceDetail(serviceID int) (*admindto.ServiceDetailResponse, *dto.ErrorResponse) {
	var service model.Service
	var shops []model.Shop

	if u.db.Find(&service, serviceID).Related(&shops).RecordNotFound() {
		return nil, dto.NotFoundError("Service")
	}

	res := &admindto.ServiceDetailResponse{}

	res.ServiceID = service.ID
	res.WifiName = service.WifiName
	res.Link = service.Link
	res.CreatedAt = service.CreatedAt
	res.UpdatedAt = service.UpdatedAt
	res.DeletedAt = service.DeletedAt
	res.ShopCount = len(shops)
	res.ShopList = []admindto.ServiceDetailResponseShopListElement{}

	for _, shop := range shops {
		reviews := u.db.Model(&model.Review{}).Where("shop_id = ?", shop.ID)
		var reviewCount int
		reviews.Count(&reviewCount)
		var average float32
		reviews.Select("avg(evaluation)").Row().Scan(&average)

		res.ShopList = append(
			res.ShopList, admindto.ServiceDetailResponseShopListElement{
				ShopID:       shop.ID,
				ShopName:     shop.ShopName,
				Area:         shop.AreaKey,
				Description:  shop.Description,
				Address:      shop.Address,
				Access:       shop.Access,
				SSID:         strings.Split(shop.SSID, ","),
				ShopType:     shop.ShopType,
				OpeningHours: shop.OpeningHours,
				SeatsNum:     shop.SeatsNum,
				HasPower:     shop.HasPower,
				ReviewCount:  reviewCount,
				Average:      average,
			},
		)
	}

	return res, nil
}

func (u *serviceUsecase) RegisterService(body *admindto.RegisterServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse) {
	service := new(model.Service)
	service.WifiName = body.WifiName
	service.Link = body.Link

	u.db.Create(&service)

	return &dto.ServiceIDResponse{
		ServiceID: service.ID,
	}, nil
}

func (u *serviceUsecase) UpdateService(serviceID int, body *admindto.UpdateServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse) {
	var service model.Service

	if u.db.Find(&service, serviceID).RecordNotFound() {
		return nil, dto.NotFoundError("Service")
	}

	if body.WifiName != "" {
		service.WifiName = body.WifiName
	}

	if body.Link != "" {
		service.Link = body.Link
	}

	u.db.Save(&service)

	return &dto.ServiceIDResponse{
		ServiceID: service.ID,
	}, nil
}

func (u *serviceUsecase) DeleteService(serviceID int) (*dto.ServiceIDResponse, *dto.ErrorResponse) {
	var service model.Service

	if u.db.Find(&service, serviceID).RecordNotFound() {
		return nil, dto.NotFoundError("Service")
	}

	u.db.Delete(&service, serviceID)

	return &dto.ServiceIDResponse{
		ServiceID: service.ID,
	}, nil
}
