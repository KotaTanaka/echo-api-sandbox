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

		if err := c.Bind(body); err != nil {
			errorResponse := data.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := data.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
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

/*
DeleteServiceAdmin | Wi-Fiサービス削除
*/
func DeleteServiceAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceIDParam := c.Param("serviceId")
		serviceID, err := strconv.Atoi(serviceIDParam)

		if err != nil {
			errorResponse := data.InvalidParameterError([]string{"ServiceID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var service model.Service

		if db.Find(&service, serviceID).RecordNotFound() {
			errorResponse := data.NotFoundError("Service")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Delete(&service, serviceID)

		return c.JSON(
			http.StatusOK,
			data.ServiceIDResponse{ServiceID: service.ID})
	}
}
