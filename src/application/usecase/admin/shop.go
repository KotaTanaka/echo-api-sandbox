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

	res := &admindto.ShopListingResponse{
		Total:    len(shops),
		ShopList: make([]admindto.ShopListingResponseElement, len(shops)),
	}

	for i, shop := range shops {
		reviewAg, err := u.reviewRepository.SelectReviewsCountAndAverageByShopID(int(shop.ID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		res.ShopList[i] = admindto.ShopListingResponseElement{
			ShopID:       shop.ID,
			ServiceID:    shop.Service.ID,
			WifiName:     shop.Service.WifiName,
			ShopName:     shop.ShopName,
			AreaKey:      shop.Area.AreaKey,
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

func (u *shopUsecase) GetShopDetail(shopID int) (*admindto.ShopDetailResponse, *dto.ErrorResponse) {
	shop, err := u.shopRepository.FindShopByID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}
	reviewAg, err := u.reviewRepository.SelectReviewsCountAndAverageByShopID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	res := &admindto.ShopDetailResponse{
		ShopID:       shop.ID,
		ServiceID:    shop.Service.ID,
		WifiName:     shop.Service.WifiName,
		ShopName:     shop.ShopName,
		AreaKey:      shop.Area.AreaKey,
		Description:  shop.Description,
		Address:      shop.Address,
		Access:       shop.Access,
		SSID:         strings.Split(shop.SSID, ","),
		ShopType:     shop.ShopType,
		OpeningHours: shop.OpeningHours,
		SeatsNum:     shop.SeatsNum,
		HasPower:     shop.HasPower,
		UpdatedAt:    shop.UpdatedAt,
		DeletedAt:    &shop.DeletedAt.Time,
		ReviewCount:  reviewAg.Count,
		ReviewList:   make([]admindto.ShopDetailResponseReviewListElement, len(shop.Reviews)),
	}

	var evaluationSum int
	for i, review := range shop.Reviews {
		evaluationSum += review.Evaluation
		res.ReviewList[i] = admindto.ShopDetailResponseReviewListElement{
			ReviewID:   review.ID,
			Comment:    review.Comment,
			Evaluation: review.Evaluation,
			Status:     review.PublishStatus,
			CreatedAt:  review.CreatedAt,
			UpdatedAt:  review.UpdatedAt,
			DeletedAt:  &review.DeletedAt.Time,
		}
	}

	if res.ReviewCount > 0 {
		res.Average = float32(evaluationSum) / float32(res.ReviewCount)
	}

	return res, nil
}

func (u *shopUsecase) RegisterShop(body *admindto.RegisterShopRequest) (*dto.ShopIDResponse, *dto.ErrorResponse) {
	shop := &model.Shop{
		ServiceID:    body.ServiceID,
		ShopName:     body.ShopName,
		AreaID:       body.AreaID,
		Description:  body.Description,
		Address:      body.Address,
		Access:       body.Access,
		SSID:         strings.Join(body.SSID, ","),
		ShopType:     body.ShopType,
		OpeningHours: body.OpeningHours,
		SeatsNum:     body.SeatsNum,
		HasPower:     body.HasPower,
	}

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
	if body.AreaID != 0 {
		shop.AreaID = body.AreaID
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
