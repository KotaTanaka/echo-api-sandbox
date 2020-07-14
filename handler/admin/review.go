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
GetReviewListAdmin | レビュー一覧取得・検索
*/
func GetReviewListAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		reviews := []model.Review{}
		db.Find(&reviews)

		response := admindata.ReviewListingResponse{}
		response.Total = len(reviews)
		response.ReviewList = []admindata.ReviewListingResponseElement{}

		for _, review := range reviews {
			shop := model.Shop{}
			db.First(&shop, review.ShopID)

			service := model.Service{}
			db.First(&service, shop.ID)

			response.ReviewList = append(
				response.ReviewList, admindata.ReviewListingResponseElement{
					ReviewID:   review.ID,
					ShopID:     shop.ID,
					ShopName:   shop.ShopName,
					ServiceID:  service.ID,
					WifiName:   service.WifiName,
					Comment:    review.Comment,
					Evaluation: review.Evaluation,
					Status:     review.PublishStatus,
					CreatedAt:  review.CreatedAt,
					UpdatedAt:  review.UpdatedAt,
					DeletedAt:  review.DeletedAt})
		}

		return c.JSON(http.StatusOK, response)
	}
}

/*
UpdateReviewStatusAdmin | レビューステータス変更
*/
func UpdateReviewStatusAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()

		reviewIDParam := c.Param("reviewId")
		reviewID, err := strconv.Atoi(reviewIDParam)

		if err != nil {
			errorResponse := data.InvalidParameterError([]string{"ReviewID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var review model.Review

		if db.Find(&review, reviewID).RecordNotFound() {
			errorResponse := data.NotFoundError("Review")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		body := new(admindata.UpdateReviewStatusRequestBody)

		if err := c.Bind(body); err != nil {
			errorResponse := data.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := data.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if body.Status == "public" {
			review.PublishStatus = true
		} else if body.Status == "hidden" {
			review.PublishStatus = false
		} else {
			errorResponse := data.InvalidParameterError([]string{"Status is 'public' or 'hidden'"})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Save(&review)

		return c.JSON(
			http.StatusOK,
			data.ReviewIDResponse{ReviewID: review.ID})
	}
}

/*
DeleteReviewAdmin | レビュー削除
*/
func DeleteReviewAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		reviewIDParam := c.Param("reviewId")
		reviewID, err := strconv.Atoi(reviewIDParam)

		if err != nil {
			errorResponse := data.InvalidParameterError([]string{"ReviewID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var review model.Review

		if db.Find(&review, reviewID).RecordNotFound() {
			errorResponse := data.NotFoundError("Review")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Delete(&review, reviewID)

		return c.JSON(
			http.StatusOK,
			data.ReviewIDResponse{ReviewID: review.ID})
	}
}
