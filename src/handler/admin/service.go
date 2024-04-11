package adminhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type ServiceHandler interface {
	GetServiceList(ctx echo.Context) error
	GetServiceDetail(ctx echo.Context) error
	RegisterService(ctx echo.Context) error
	UpdateService(ctx echo.Context) error
	DeleteService(ctx echo.Context) error
}

type serviceHandler struct {
	db *gorm.DB
}

func NewServiceHandler(db *gorm.DB) ServiceHandler {
	return &serviceHandler{db: db}
}

func (sh serviceHandler) GetServiceList(ctx echo.Context) error {
	services := []model.Service{}
	sh.db.Find(&services)

	response := admindto.ServiceListingResponse{}
	response.Total = len(services)
	response.ServiceList = []admindto.ServiceListingResponseElement{}

	for _, service := range services {
		shopCount := 0
		sh.db.Model(&model.Shop{}).Where("service_id = ?", service.ID).Count(&shopCount)

		response.ServiceList = append(
			response.ServiceList, admindto.ServiceListingResponseElement{
				ServiceID: service.ID,
				WifiName:  service.WifiName,
				Link:      service.Link,
				ShopCount: shopCount,
			},
		)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (sh serviceHandler) GetServiceDetail(ctx echo.Context) error {
	serviceIDParam := ctx.Param("serviceId")
	serviceID, err := strconv.Atoi(serviceIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ServiceID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var service model.Service
	var shops []model.Shop

	if sh.db.Find(&service, serviceID).Related(&shops).RecordNotFound() {
		errorResponse := dto.NotFoundError("Service")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
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
		reviews := sh.db.Model(&model.Review{}).Where("shop_id = ?", shop.ID)
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
				Average:      average,
			},
		)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (sh serviceHandler) RegisterService(ctx echo.Context) error {
	validator.New()
	body := new(admindto.RegisterServiceRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	service := new(model.Service)
	service.WifiName = body.WifiName
	service.Link = body.Link

	sh.db.Create(&service)

	return ctx.JSON(
		http.StatusOK,
		dto.ServiceIDResponse{
			ServiceID: service.ID,
		},
	)
}

func (sh serviceHandler) UpdateService(ctx echo.Context) error {
	validator.New()

	serviceIDParam := ctx.Param("serviceId")
	serviceID, err := strconv.Atoi(serviceIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ServiceID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var service model.Service

	if sh.db.Find(&service, serviceID).RecordNotFound() {
		errorResponse := dto.NotFoundError("Service")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	body := new(admindto.UpdateServiceRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if body.WifiName != "" {
		service.WifiName = body.WifiName
	}

	if body.Link != "" {
		service.Link = body.Link
	}

	sh.db.Save(&service)

	return ctx.JSON(
		http.StatusOK,
		dto.ServiceIDResponse{
			ServiceID: service.ID,
		},
	)
}

func (sh serviceHandler) DeleteService(ctx echo.Context) error {
	serviceIDParam := ctx.Param("serviceId")
	serviceID, err := strconv.Atoi(serviceIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ServiceID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var service model.Service

	if sh.db.Find(&service, serviceID).RecordNotFound() {
		errorResponse := dto.NotFoundError("Service")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	sh.db.Delete(&service, serviceID)

	return ctx.JSON(
		http.StatusOK,
		dto.ServiceIDResponse{
			ServiceID: service.ID,
		},
	)
}
