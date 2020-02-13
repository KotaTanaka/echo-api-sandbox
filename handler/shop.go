/*
Package handler | Wi-Fi提供店舗関連ハンドラー
*/
package handler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"../data"
)

/*
GetShopListClient | 店舗一覧取得
*/
func GetShopListClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shops := []data.Shop{}
		db.Find(&shops)

		response := data.ShopListingResponse{}
		response.Total = len(shops)

		for _, shop := range shops {
			service := data.Service{}
			db.First(&service, shop.ServiceID)

			response.ShopList = append(
				// TODO SSID: 文字列を配列に変換
				// TODO Average: 評価の平均値の計算
				response.ShopList, data.ShopListingResponseElement{
					ShopID:       shop.ID,
					ShopName:     shop.ShopName,
					WifiName:     service.WifiName,
					ServiceLink:  service.Link,
					Ssid:         []string{shop.SSID},
					Address:      shop.Address,
					Acceess:      "",
					Description:  shop.Description,
					ShopType:     shop.ShopType,
					OpeningHours: shop.OpeningHours,
					SeatsNum:     shop.SeatsNum,
					Power:        shop.HasPower,
					ReviewCount:  len(shop.Reviews),
					Average:      0})
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
		body := new(data.RegisterShopRequestBody)
		errorResponse := new(data.ErrorResponse)

		if err := c.Bind(body); err != nil {
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Invalid Request"
			errorResponse.DetailMessage = []string{err.Error()}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Invalid Parameter"
			errorResponse.DetailMessage = strings.Split(err.(validator.ValidationErrors).Error(), "\n")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		shop := new(data.Shop)
		shop.ServiceID = body.ServiceID
		shop.SSID = body.SSID
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
