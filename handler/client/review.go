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
