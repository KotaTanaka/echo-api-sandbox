/*
Package adminhandler 管理API関連ハンドラー
*/
package adminhandler

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

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
