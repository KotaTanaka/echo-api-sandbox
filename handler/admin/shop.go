/*
Package adminhandler 管理API関連ハンドラー
*/
package adminhandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"../../data"
	admindata "../../data/admin"
	"../../model"
)

/*
RegisterShopAdmin | 店舗登録
*/
func RegisterShopAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindata.RegisterShopRequestBody)
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

		shop := new(model.Shop)
		shop.ServiceID = body.ServiceID
		shop.AreaKey = body.Area
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
