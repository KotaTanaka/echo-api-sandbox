/*
Package adminhandler 管理API関連ハンドラー
*/
package adminhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"../../data"
	admindata "../../data/admin"
	"../../model"
)

/*
GetShopListAdmin | 店舗一覧取得・検索
*/
func GetShopListAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shops := []model.Shop{}
		db.Find(&shops)

		response := admindata.ShopListingResponse{}
		response.Total = len(shops)

		for _, shop := range shops {
			service := model.Service{}
			db.First(&service, shop.ServiceID)

			reviews := db.Model(&model.Review{}).Where("shop_id = ?", shop.ID)
			var reviewCount int
			reviews.Count(&reviewCount)
			var average float32
			reviews.Select("avg(evaluation)").Row().Scan(&average)

			response.ShopList = append(
				response.ShopList, admindata.ShopListingResponseElement{
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
					ReviewCount:  reviewCount,
					Average:      average})
		}

		return c.JSON(http.StatusOK, response)
	}
}

/*
RegisterShopAdmin | 店舗登録
*/
func RegisterShopAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindata.RegisterShopRequestBody)

		if err := c.Bind(body); err != nil {
			errorResponse := data.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := data.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		shop := new(model.Shop)
		shop.ServiceID = body.ServiceID
		shop.AreaKey = body.Area
		shop.SSID = strings.Join(body.SSID, ",")
		shop.ShopName = body.ShopName
		shop.Description = body.Description
		shop.Address = body.Address
		shop.ShopType = body.ShopType
		shop.OpeningHours = body.OpeningHours
		shop.SeatsNum = body.SeatsNum
		shop.HasPower = body.HasPower

		db.Create(&shop)

		return c.JSON(
			http.StatusOK,
			data.ShopIDResponse{ShopID: shop.ID})
	}
}

/*
DeleteShopAdmin | 店舗削除
*/
func DeleteShopAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shopIDParam := c.Param("shopId")
		shopID, err := strconv.Atoi(shopIDParam)

		if err != nil {
			errorResponse := data.InvalidParameterError([]string{"ShopID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var shop model.Shop

		if db.Find(&shop, shopID).RecordNotFound() {
			errorResponse := data.NotFoundError("Shop")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Delete(&shop, shopID)

		return c.JSON(
			http.StatusOK,
			data.ShopIDResponse{ShopID: shop.ID})
	}
}
