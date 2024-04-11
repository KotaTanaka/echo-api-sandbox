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

type ShopHandler interface {
	GetShopList(ctx echo.Context) error
	GetShopDetail(ctx echo.Context) error
	RegisterShop(ctx echo.Context) error
	UpdateShop(ctx echo.Context) error
	DeleteShop(ctx echo.Context) error
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

	response := admindto.ShopListingResponse{}
	response.Total = len(shops)
	response.ShopList = []admindto.ShopListingResponseElement{}

	for _, shop := range shops {
		service := model.Service{}
		sh.db.First(&service, shop.ServiceID)

		reviews := sh.db.Model(&model.Review{}).Where("shop_id = ?", shop.ID)
		var reviewCount int
		reviews.Count(&reviewCount)
		var average float32
		reviews.Select("avg(evaluation)").Row().Scan(&average)

		response.ShopList = append(
			response.ShopList, admindto.ShopListingResponseElement{
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

	return ctx.JSON(http.StatusOK, response)
}

func (sh shopHandler) GetShopDetail(ctx echo.Context) error {
	shopIDParam := ctx.Param("shopId")
	shopID, err := strconv.Atoi(shopIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ShopID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var shop model.Shop
	var service model.Service
	var reviews []model.Review

	if sh.db.Find(&shop, shopID).Related(&reviews).RecordNotFound() {
		errorResponse := dto.NotFoundError("Shop")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	sh.db.First(&service, shop.ServiceID)

	response := admindto.ShopDetailResponse{}

	response.ShopID = shop.ID
	response.ServiceID = service.ID
	response.WifiName = service.WifiName
	response.ShopName = shop.ShopName
	response.Area = shop.AreaKey
	response.Description = shop.Description
	response.Address = shop.Address
	response.Access = shop.Access
	response.SSID = strings.Split(shop.SSID, ",")
	response.ShopType = shop.ShopType
	response.OpeningHours = shop.OpeningHours
	response.SeatsNum = shop.SeatsNum
	response.SeatsNum = shop.SeatsNum
	response.HasPower = shop.HasPower
	response.UpdatedAt = shop.UpdatedAt
	response.DeletedAt = shop.DeletedAt
	response.ReviewCount = len(reviews)
	response.ReviewList = []admindto.ShopDetailResponseReviewListElement{}

	var evaluationSum int
	for _, review := range reviews {
		evaluationSum += review.Evaluation
		response.ReviewList = append(
			response.ReviewList, admindto.ShopDetailResponseReviewListElement{
				ReviewID:   review.ID,
				Comment:    review.Comment,
				Evaluation: review.Evaluation,
				Status:     review.PublishStatus,
				CreatedAt:  review.CreatedAt,
				UpdatedAt:  review.UpdatedAt,
				DeletedAt:  review.DeletedAt,
			})
	}

	if response.ReviewCount > 0 {
		response.Average = float32(evaluationSum) / float32(response.ReviewCount)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (sh shopHandler) RegisterShop(ctx echo.Context) error {
	validator.New()
	body := new(admindto.RegisterShopRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	shop := new(model.Shop)
	shop.ServiceID = body.ServiceID
	shop.ShopName = body.ShopName
	shop.AreaKey = body.Area
	shop.Description = body.Description
	shop.Address = body.Address
	shop.Access = body.Access
	shop.SSID = strings.Join(body.SSID, ",")
	shop.ShopType = body.ShopType
	shop.OpeningHours = body.OpeningHours
	shop.SeatsNum = body.SeatsNum
	shop.HasPower = body.HasPower

	sh.db.Create(&shop)

	return ctx.JSON(
		http.StatusOK,
		dto.ShopIDResponse{
			ShopID: shop.ID,
		},
	)
}

func (sh shopHandler) UpdateShop(ctx echo.Context) error {
	validator.New()

	shopIDParam := ctx.Param("shopId")
	shopID, err := strconv.Atoi(shopIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ShopID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var shop model.Shop

	if sh.db.Find(&shop, shopID).RecordNotFound() {
		errorResponse := dto.NotFoundError("Shop")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	body := new(admindto.UpdateShopRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if body.ShopName != "" {
		shop.ShopName = body.ShopName
	}
	if body.Area != "" {
		shop.AreaKey = body.Area
	}
	if body.Description != "" {
		shop.Description = body.Description
	}
	if body.Address != "" {
		shop.Address = body.Address
	}
	if body.Access != "" {
		shop.Access = body.Access
	}
	if len(body.SSID) > 0 {
		shop.SSID = strings.Join(body.SSID, ",")
	}
	if body.ShopType != "" {
		shop.ShopType = body.ShopType
	}
	if body.OpeningHours != "" {
		shop.OpeningHours = body.OpeningHours
	}
	if body.SeatsNum != 0 {
		shop.SeatsNum = body.SeatsNum
	}
	if body.HasPower {
		shop.HasPower = body.HasPower
	}

	sh.db.Save(&shop)

	return ctx.JSON(
		http.StatusOK,
		dto.ShopIDResponse{
			ShopID: shop.ID,
		},
	)
}

func (sh shopHandler) DeleteShop(ctx echo.Context) error {
	shopIDParam := ctx.Param("shopId")
	shopID, err := strconv.Atoi(shopIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ShopID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var shop model.Shop

	if sh.db.Find(&shop, shopID).RecordNotFound() {
		errorResponse := dto.NotFoundError("Shop")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	sh.db.Delete(&shop, shopID)

	return ctx.JSON(
		http.StatusOK,
		dto.ShopIDResponse{
			ShopID: shop.ID,
		},
	)
}
