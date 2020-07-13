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
			shopCount := 0
			db.Model(&model.Shop{}).Where("service_id = ?", service.ID).Count(&shopCount)

			response.ServiceList = append(
				response.ServiceList, admindata.ServiceListingResponseElement{
					ServiceID: service.ID,
					WifiName:  service.WifiName,
					Link:      service.Link,
					ShopCount: shopCount})
		}

		return c.JSON(http.StatusOK, response)
	}
}

/*
GetServiceDetailAdmin | Wi-Fiサービス詳細取得
*/
func GetServiceDetailAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceIDParam := c.Param("serviceId")
		serviceID, err := strconv.Atoi(serviceIDParam)

		if err != nil {
			errorResponse := data.InvalidParameterError([]string{"ServiceID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var service model.Service
		var shops []model.Shop

		if db.Find(&service, serviceID).Related(&shops).RecordNotFound() {
			errorResponse := data.NotFoundError("Service")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		response := admindata.ServiceDetailResponse{}

		response.ServiceID = service.ID
		response.WifiName = service.WifiName
		response.Link = service.Link
		response.CreatedAt = service.CreatedAt
		response.UpdatedAt = service.UpdatedAt
		response.DeletedAt = service.DeletedAt
		response.ShopCount = len(shops)

		for _, shop := range shops {
			// TODO SSID: 文字列を配列に変換
			// TODO Average: 評価の平均値の計算
			response.ShopList = append(
				response.ShopList, admindata.ServiceDetailResponseShopListElement{
					ShopID:       shop.ID,
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
