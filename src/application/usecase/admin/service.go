package adminusecase

import (
	"strings"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type ServiceUsecase interface {
	GetServiceList() (*admindto.ServiceListingResponse, *dto.ErrorResponse)
	GetServiceDetail(serviceID int) (*admindto.ServiceDetailResponse, *dto.ErrorResponse)
	RegisterService(body *admindto.RegisterServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse)
	UpdateService(serviceID int, body *admindto.UpdateServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse)
	DeleteService(serviceID int) (*dto.ServiceIDResponse, *dto.ErrorResponse)
}

type serviceUsecase struct {
	serviceRepository repository.ServiceRepository
	shopRepository    repository.ShopRepository
	reviewRepository  repository.ReviewRepository
}

func NewServiceUsecase(
	serviceRepository repository.ServiceRepository,
	shopRepository repository.ShopRepository,
	reviewRepository repository.ReviewRepository,
) ServiceUsecase {
	return &serviceUsecase{
		serviceRepository: serviceRepository,
		shopRepository:    shopRepository,
		reviewRepository:  reviewRepository,
	}
}

func (u *serviceUsecase) GetServiceList() (*admindto.ServiceListingResponse, *dto.ErrorResponse) {
	services, err := u.serviceRepository.ListServices()
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	res := &admindto.ServiceListingResponse{
		Total:       len(services),
		ServiceList: make([]admindto.ServiceListingResponseElement, len(services)),
	}

	for i, service := range services {
		shopAg, err := u.shopRepository.CountShopsByServiceID(int(service.ID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		res.ServiceList[i] = admindto.ServiceListingResponseElement{
			ServiceID: service.ID,
			WifiName:  service.WifiName,
			Link:      service.Link,
			ShopCount: shopAg.Count,
		}
	}

	return res, nil
}

func (u *serviceUsecase) GetServiceDetail(serviceID int) (*admindto.ServiceDetailResponse, *dto.ErrorResponse) {
	service, err := u.serviceRepository.FindServiceByID(serviceID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	shops, err := u.shopRepository.ListShopsByServiceID(serviceID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	res := &admindto.ServiceDetailResponse{
		ServiceID: service.ID,
		WifiName:  service.WifiName,
		Link:      service.Link,
		CreatedAt: service.CreatedAt,
		UpdatedAt: service.UpdatedAt,
		DeletedAt: service.DeletedAt,
		ShopCount: len(shops),
		ShopList:  make([]admindto.ServiceDetailResponseShopListElement, len(shops)),
	}

	for i, shop := range shops {
		reviewAg, err := u.reviewRepository.SelectReviewsCountAndAverageByShopID(int(shop.ID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		res.ShopList[i] = admindto.ServiceDetailResponseShopListElement{
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
			ReviewCount:  reviewAg.Count,
			Average:      reviewAg.Average,
		}
	}

	return res, nil
}

func (u *serviceUsecase) RegisterService(body *admindto.RegisterServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse) {
	service := &model.Service{
		WifiName: body.WifiName,
		Link:     body.Link,
	}

	service, err := u.serviceRepository.CreateService(service)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ServiceIDResponse{
		ServiceID: service.ID,
	}, nil
}

func (u *serviceUsecase) UpdateService(serviceID int, body *admindto.UpdateServiceRequest) (*dto.ServiceIDResponse, *dto.ErrorResponse) {
	service, err := u.serviceRepository.FindServiceByID(serviceID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	if body.WifiName != "" {
		service.WifiName = body.WifiName
	}

	if body.Link != "" {
		service.Link = body.Link
	}

	service, err = u.serviceRepository.UpdateService(service)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ServiceIDResponse{
		ServiceID: service.ID,
	}, nil
}

func (u *serviceUsecase) DeleteService(serviceID int) (*dto.ServiceIDResponse, *dto.ErrorResponse) {
	service, err := u.serviceRepository.FindServiceByID(serviceID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	err = u.serviceRepository.DeleteService(service)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ServiceIDResponse{
		ServiceID: service.ID,
	}, nil
}
