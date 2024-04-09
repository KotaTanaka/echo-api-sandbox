package clienthandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientdto "github.com/KotaTanaka/echo-api-sandbox/model/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/model/entity"
)

func GetShopListClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shops := []entity.Shop{}
		db.Find(&shops)

		response := clientdto.ShopListingResponse{}
		response.Total = len(shops)
		response.ShopList = []clientdto.ShopListingResponseElement{}

		for _, shop := range shops {
			service := entity.Service{}
			db.First(&service, shop.ServiceID)

			reviews := db.Model(&entity.Review{}).Where("shop_id = ?", shop.ID)
			var reviewCount int
			reviews.Count(&reviewCount)
			var average float32
			reviews.Select("avg(evaluation)").Row().Scan(&average)

			response.ShopList = append(
				response.ShopList, clientdto.ShopListingResponseElement{
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
					Average:      average})
		}

		return c.JSON(http.StatusOK, response)
	}
}
