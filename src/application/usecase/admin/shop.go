package adminusecase

import (
	"strings"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type ShopUsecase interface {
	GetShopList() (*admindto.ShopListingResponse, *dto.ErrorResponse)
	GetShopDetail(shopID int) (*admindto.ShopDetailResponse, *dto.ErrorResponse)
	RegisterShop(body *admindto.RegisterShopRequest) (*dto.ShopIDResponse, *dto.ErrorResponse)
	UpdateShop(shopID int, body *admindto.UpdateShopRequest) (*dto.ShopIDResponse, *dto.ErrorResponse)
	DeleteShop(shopID int) (*dto.ShopIDResponse, *dto.ErrorResponse)
}

type shopUsecase struct {
	serviceRepository repository.ServiceRepository
	shopRepository    repository.ShopRepository
	reviewRepository  repository.ReviewRepository
}

func NewShopUsecase(
	serviceRepository repository.ServiceRepository,
	shopRepository repository.ShopRepository,
	reviewRepository repository.ReviewRepository,
) ShopUsecase {
	return &shopUsecase{
		serviceRepository: serviceRepository,
		shopRepository:    shopRepository,
		reviewRepository:  reviewRepository,
	}
}

func (u *shopUsecase) GetShopList() (*admindto.ShopListingResponse, *dto.ErrorResponse) {
	shops, err := u.shopRepository.ListShops()
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	res := &admindto.ShopListingResponse{}
	res.Total = len(shops)
	res.ShopList = []admindto.ShopListingResponseElement{}

	for _, shop := range shops {
		service, err := u.serviceRepository.FindServiceByID(int(shop.ServiceID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		reviewAg, err := u.reviewRepository.SelectReviewsCountAndAverageByShopID(int(shop.ID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		res.ShopList = append(
			res.ShopList, admindto.ShopListingResponseElement{
				ShopID:       shop.ID,
				ServiceID:    service.ID,
				WifiName:     service.WifiName,
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
			},
		)
	}

	return res, nil
}

func (u *shopUsecase) GetShopDetail(shopID int) (*admindto.ShopDetailResponse, *dto.ErrorResponse) {
	shop, err := u.shopRepository.FindShopByID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	service, err := u.serviceRepository.FindServiceByID(int(shop.ServiceID))
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	reviews, err := u.reviewRepository.ListReviewsByShopID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	res := &admindto.ShopDetailResponse{}

	res.ShopID = shop.ID
	res.ServiceID = service.ID
	res.WifiName = service.WifiName
	res.ShopName = shop.ShopName
	res.Area = shop.AreaKey
	res.Description = shop.Description
	res.Address = shop.Address
	res.Access = shop.Access
	res.SSID = strings.Split(shop.SSID, ",")
	res.ShopType = shop.ShopType
	res.OpeningHours = shop.OpeningHours
	res.SeatsNum = shop.SeatsNum
	res.SeatsNum = shop.SeatsNum
	res.HasPower = shop.HasPower
	res.UpdatedAt = shop.UpdatedAt
	res.DeletedAt = shop.DeletedAt
	res.ReviewCount = len(reviews)
	res.ReviewList = []admindto.ShopDetailResponseReviewListElement{}

	var evaluationSum int
	for _, review := range reviews {
		evaluationSum += review.Evaluation
		res.ReviewList = append(
			res.ReviewList, admindto.ShopDetailResponseReviewListElement{
				ReviewID:   review.ID,
				Comment:    review.Comment,
				Evaluation: review.Evaluation,
				Status:     review.PublishStatus,
				CreatedAt:  review.CreatedAt,
				UpdatedAt:  review.UpdatedAt,
				DeletedAt:  review.DeletedAt,
			})
	}

	if res.ReviewCount > 0 {
		res.Average = float32(evaluationSum) / float32(res.ReviewCount)
	}

	return res, nil
}

func (u *shopUsecase) RegisterShop(body *admindto.RegisterShopRequest) (*dto.ShopIDResponse, *dto.ErrorResponse) {
	shop := new(model.Shop)

	shop.ServiceID = body.ServiceID
	shop.ShopName = body.ShopName
	shop.AreaKey = body.Area
	shop.Description = body.Description
	shop.Address = body.Address
	shop.Access = body.Access
	shop.SSID = strings.Join(body.SSID, ",")
	shop.ShopType = body.ShopType
	shop.OpeningHours = body.OpeningHours
	shop.SeatsNum = body.SeatsNum
	shop.HasPower = body.HasPower

	shop, err := u.shopRepository.CreateShop(shop)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ShopIDResponse{
		ShopID: shop.ID,
	}, nil
}

func (u *shopUsecase) UpdateShop(shopID int, body *admindto.UpdateShopRequest) (*dto.ShopIDResponse, *dto.ErrorResponse) {
	shop, err := u.shopRepository.FindShopByID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	if body.ShopName != "" {
		shop.ShopName = body.ShopName
	}
	if body.Area != "" {
		shop.AreaKey = body.Area
	}
	if body.Description != "" {
		shop.Description = body.Description
	}
	if body.Address != "" {
		shop.Address = body.Address
	}
	if body.Access != "" {
		shop.Access = body.Access
	}
	if len(body.SSID) > 0 {
		shop.SSID = strings.Join(body.SSID, ",")
	}
	if body.ShopType != "" {
		shop.ShopType = body.ShopType
	}
	if body.OpeningHours != "" {
		shop.OpeningHours = body.OpeningHours
	}
	if body.SeatsNum != 0 {
		shop.SeatsNum = body.SeatsNum
	}
	if body.HasPower {
		shop.HasPower = body.HasPower
	}

	shop, err = u.shopRepository.UpdateShop(shop)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ShopIDResponse{
		ShopID: shop.ID,
	}, nil
}

func (u *shopUsecase) DeleteShop(shopID int) (*dto.ShopIDResponse, *dto.ErrorResponse) {
	shop, err := u.shopRepository.FindShopByID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	err = u.shopRepository.DeleteShop(shop)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ShopIDResponse{
		ShopID: shop.ID,
	}, nil
}
