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
GetServiceListAdmin | Wi-Fiサービス一覧取得・検索
*/
func GetServiceListAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		services := []model.Service{}
		db.Find(&services)

		response := admindata.ServiceListingResponse{}
		response.Total = len(services)

		for _, service := range services {
			response.ServiceList = append(
				response.ServiceList, admindata.ServiceListingResponseElement{
					ServiceID: service.ID,
					WifiName:  service.WifiName,
					Link:      service.Link,
					ShopCount: len(service.Shops)})
		}

		return c.JSON(http.StatusOK, response)
	}
}

/*
RegisterServiceAdmin | Wi-Fiサービス登録
*/
func RegisterServiceAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindata.RegisterServiceRequestBody)
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

		service := new(model.Service)
		service.WifiName = body.WifiName
		service.Link = body.Link

		db.Create(&service)

		return c.JSON(
			http.StatusOK,
			data.ServiceIDResponse{ServiceID: service.ID})
	}
}
