/*
Package clienthandler クライアントAPI関連ハンドラー
*/
package clienthandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"../../data"
	clientdata "../../data/client"
	"../../model"
)

/*
GetReviewListClient | 店舗に紐づくレビュー一覧取得
*/
func GetReviewListClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shopID := c.QueryParam("shopId")

		var shop model.Shop
		var service model.Service
		var reviews []model.Review

		if db.Where("id = ?", shopID).Find(&shop).Related(&reviews).RecordNotFound() {
			errorResponse := new(data.ErrorResponse)
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Invalid Request"
			errorResponse.DetailMessage = []string{"Shop Not Found."}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.First(&service, shop.ID)

		response := clientdata.ReviewListingResponse{}
		response.ShopID = shop.ID
		response.ShopName = shop.ShopName
		response.ServiceID = service.ID
		response.WifiName = service.WifiName
		// TODO レビューの平均値算出
		response.Average = 0
		response.Total = len(reviews)

		for _, review := range reviews {
			response.ReviewList = append(
				response.ReviewList, clientdata.ReviewListingResponseElement{
					ReviewID:   review.ID,
					Comment:    review.Comment,
					Evaluation: review.Evaluation,
					Status:     review.PublishStatus,
					CreatedAt:  review.CreatedAt})
		}

		return c.JSON(http.StatusOK, response)
	}
}

/*
CreateReviewClient | 店舗へのレビュー投稿
*/
func CreateReviewClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(clientdata.CreateReviewRequestBody)
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

		shop := model.Shop{}
		db.First(&shop, body.ShopID)

		review := new(model.Review)
		review.ShopID = shop.ID
		review.Comment = body.Comment
		review.Evaluation = body.Evaluation

		db.Create(&review)

		return c.JSON(
			http.StatusOK,
			data.ReviewIDResponse{ReviewID: review.ID})
	}
}
