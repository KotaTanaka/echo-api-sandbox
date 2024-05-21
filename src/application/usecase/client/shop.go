package clientusecase

import (
	"fmt"
	"strings"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type ShopUsecase interface {
	GetShopList() (*clientdto.ShopListingResponse, *dto.ErrorResponse)
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

func (u *shopUsecase) GetShopList() (*clientdto.ShopListingResponse, *dto.ErrorResponse) {
	shops, err := u.shopRepository.ListShops()
	if err != nil {
		return nil, dto.HandleDBError(err, "Shops")
	}

	res := &clientdto.ShopListingResponse{
		Total:    len(shops),
		ShopList: make([]clientdto.ShopListingResponseElement, len(shops)),
	}

	for i, shop := range shops {
		reviewAg, err := u.reviewRepository.SelectReviewsCountAndAverageByShopID(int(shop.ID))
		if err != nil {
			return nil, dto.HandleDBError(err, fmt.Sprintf("Reviews(Shop ID:%d)", shop.ID))
		}

		res.ShopList[i] = clientdto.ShopListingResponseElement{
			ShopID:       shop.ID,
			WifiName:     shop.Service.WifiName,
			ServiceLink:  shop.Service.Link,
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
