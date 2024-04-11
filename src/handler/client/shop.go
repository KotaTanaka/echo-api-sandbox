package clienthandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type ShopHandler interface {
	GetShopList(ctx echo.Context) error
}

type shopHandler struct {
	db *gorm.DB
}

func NewShopHandler(db *gorm.DB) ShopHandler {
	return &shopHandler{db: db}
}

func (sh shopHandler) GetShopList(ctx echo.Context) error {
	shops := []model.Shop{}
	sh.db.Find(&shops)

	response := clientdto.ShopListingResponse{}
	response.Total = len(shops)
	response.ShopList = []clientdto.ShopListingResponseElement{}

	for _, shop := range shops {
		service := model.Service{}
		sh.db.First(&service, shop.ServiceID)

		reviews := sh.db.Model(&model.Review{}).Where("shop_id = ?", shop.ID)
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

	return ctx.JSON(http.StatusOK, response)
}
