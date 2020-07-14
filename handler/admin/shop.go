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
		response.ShopList = []admindata.ShopListingResponseElement{}

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
GetShopDetailAdmin | 店舗詳細取得
*/
func GetShopDetailAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shopIDParam := c.Param("shopId")
		shopID, err := strconv.Atoi(shopIDParam)

		if err != nil {
			errorResponse := data.InvalidParameterError([]string{"ShopID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var shop model.Shop
		var service model.Service
		var reviews []model.Review

		if db.Find(&shop, shopID).Related(&reviews).RecordNotFound() {
			errorResponse := data.NotFoundError("Shop")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.First(&service, shop.ServiceID)

		response := admindata.ShopDetailResponse{}

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
		response.ReviewList = []admindata.ShopDetailResponseReviewListElement{}

		var evaluationSum int
		for _, review := range reviews {
			evaluationSum += review.Evaluation
			response.ReviewList = append(
				response.ReviewList, admindata.ShopDetailResponseReviewListElement{
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

		db.Create(&shop)

		return c.JSON(
			http.StatusOK,
			data.ShopIDResponse{ShopID: shop.ID})
	}
}

/*
UpdateShopAdmin | 店舗編集
*/
func UpdateShopAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()

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

		body := new(admindata.UpdateShopRequestBody)

		if err := c.Bind(body); err != nil {
			errorResponse := data.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := data.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
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
		if body.HasPower != false {
			shop.HasPower = body.HasPower
		}

		db.Save(&shop)

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
