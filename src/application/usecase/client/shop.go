package clientusecase

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type ShopUsecase interface {
	GetShopList() (*clientdto.ShopListingResponse, *dto.ErrorResponse)
}

type shopUsecase struct {
	db *gorm.DB
}

func NewShopUsecase(db *gorm.DB) ShopUsecase {
	return &shopUsecase{db: db}
}

func (u *shopUsecase) GetShopList() (*clientdto.ShopListingResponse, *dto.ErrorResponse) {
	shops := []model.Shop{}
	u.db.Find(&shops)

	res := &clientdto.ShopListingResponse{}
	res.Total = len(shops)
	res.ShopList = []clientdto.ShopListingResponseElement{}

	for _, shop := range shops {
		service := model.Service{}
		u.db.First(&service, shop.ServiceID)

		reviews := u.db.Model(&model.Review{}).Where("shop_id = ?", shop.ID)
		var reviewCount int
		reviews.Count(&reviewCount)
		var average float32
		reviews.Select("avg(evaluation)").Row().Scan(&average)

		res.ShopList = append(
			res.ShopList, clientdto.ShopListingResponseElement{
				ShopID:       shop.ID,
				WifiName:     service.WifiName,
				ServiceLink:  service.Link,
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
