package adminhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/model/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/model/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/model/entity"
)

func GetServiceListAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		services := []entity.Service{}
		db.Find(&services)

		response := admindto.ServiceListingResponse{}
		response.Total = len(services)
		response.ServiceList = []admindto.ServiceListingResponseElement{}

		for _, service := range services {
			shopCount := 0
			db.Model(&entity.Shop{}).Where("service_id = ?", service.ID).Count(&shopCount)

			response.ServiceList = append(
				response.ServiceList, admindto.ServiceListingResponseElement{
					ServiceID: service.ID,
					WifiName:  service.WifiName,
					Link:      service.Link,
					ShopCount: shopCount})
		}

		return c.JSON(http.StatusOK, response)
	}
}

func GetServiceDetailAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceIDParam := c.Param("serviceId")
		serviceID, err := strconv.Atoi(serviceIDParam)

		if err != nil {
			errorResponse := dto.InvalidParameterError([]string{"ServiceID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var service entity.Service
		var shops []entity.Shop

		if db.Find(&service, serviceID).Related(&shops).RecordNotFound() {
			errorResponse := dto.NotFoundError("Service")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		response := admindto.ServiceDetailResponse{}

		response.ServiceID = service.ID
		response.WifiName = service.WifiName
		response.Link = service.Link
		response.CreatedAt = service.CreatedAt
		response.UpdatedAt = service.UpdatedAt
		response.DeletedAt = service.DeletedAt
		response.ShopCount = len(shops)
		response.ShopList = []admindto.ServiceDetailResponseShopListElement{}

		for _, shop := range shops {
			reviews := db.Model(&entity.Review{}).Where("shop_id = ?", shop.ID)
			var reviewCount int
			reviews.Count(&reviewCount)
			var average float32
			reviews.Select("avg(evaluation)").Row().Scan(&average)

			response.ShopList = append(
				response.ShopList, admindto.ServiceDetailResponseShopListElement{
					ShopID:       shop.ID,
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

func RegisterServiceAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindto.RegisterServiceRequest)

		if err := c.Bind(body); err != nil {
			errorResponse := dto.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		service := new(entity.Service)
		service.WifiName = body.WifiName
		service.Link = body.Link

		db.Create(&service)

		return c.JSON(
			http.StatusOK,
			dto.ServiceIDResponse{ServiceID: service.ID})
	}
}

func UpdateServiceAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()

		serviceIDParam := c.Param("serviceId")
		serviceID, err := strconv.Atoi(serviceIDParam)

		if err != nil {
			errorResponse := dto.InvalidParameterError([]string{"ServiceID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var service entity.Service

		if db.Find(&service, serviceID).RecordNotFound() {
			errorResponse := dto.NotFoundError("Service")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		body := new(admindto.UpdateServiceRequest)

		if err := c.Bind(body); err != nil {
			errorResponse := dto.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if body.WifiName != "" {
			service.WifiName = body.WifiName
		}

		if body.Link != "" {
			service.Link = body.Link
		}

		db.Save(&service)

		return c.JSON(
			http.StatusOK,
			dto.ServiceIDResponse{ServiceID: service.ID})
	}
}

func DeleteServiceAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceIDParam := c.Param("serviceId")
		serviceID, err := strconv.Atoi(serviceIDParam)

		if err != nil {
			errorResponse := dto.InvalidParameterError([]string{"ServiceID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var service entity.Service

		if db.Find(&service, serviceID).RecordNotFound() {
			errorResponse := dto.NotFoundError("Service")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Delete(&service, serviceID)

		return c.JSON(
			http.StatusOK,
			dto.ServiceIDResponse{ServiceID: service.ID})
	}
}
