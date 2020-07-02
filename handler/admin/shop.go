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

			response.ShopList = append(
				// TODO SSID: 文字列を配列に変換
				// TODO Average: 評価の平均値の計算
				response.ShopList, admindata.ShopListingResponseElement{
					ShopID:       shop.ID,
					ServiceID:    service.ID,
					WifiName:     service.WifiName,
					ShopName:     shop.ShopName,
					Area:         shop.AreaKey,
					Description:  shop.Description,
					Address:      shop.Address,
					Access:       shop.Access,
					SSID:         []string{shop.SSID},
					ShopType:     shop.ShopType,
					OpeningHours: shop.OpeningHours,
					SeatsNum:     shop.SeatsNum,
					HasPower:     shop.HasPower,
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
