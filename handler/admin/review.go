/*
Package adminhandler 管理API関連ハンドラー
*/
package adminhandler

import (
	"net/http"

	admindata "../../data/admin"
	"../../model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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

		for _, review := range reviews {
			shop := model.Shop{}
			db.First(&shop, review.ShopID)

			service := model.Service{}
			db.First(&service, shop.ID)

			response.ReviewList = append(
				response.ReviewList, admindata.ReviewListingResponseElement{
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
