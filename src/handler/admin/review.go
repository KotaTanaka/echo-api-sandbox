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

func GetReviewListAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		reviews := []entity.Review{}
		db.Find(&reviews)

		response := admindto.ReviewListingResponse{}
		response.Total = len(reviews)
		response.ReviewList = []admindto.ReviewListingResponseElement{}

		for _, review := range reviews {
			shop := entity.Shop{}
			db.First(&shop, review.ShopID)

			service := entity.Service{}
			db.First(&service, shop.ID)

			response.ReviewList = append(
				response.ReviewList, admindto.ReviewListingResponseElement{
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

func UpdateReviewStatusAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()

		reviewIDParam := c.Param("reviewId")
		reviewID, err := strconv.Atoi(reviewIDParam)

		if err != nil {
			errorResponse := dto.InvalidParameterError([]string{"ReviewID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var review entity.Review

		if db.Find(&review, reviewID).RecordNotFound() {
			errorResponse := dto.NotFoundError("Review")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		body := new(admindto.UpdateReviewStatusRequest)

		if err := c.Bind(body); err != nil {
			errorResponse := dto.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if body.Status == "public" {
			review.PublishStatus = true
		} else if body.Status == "hidden" {
			review.PublishStatus = false
		} else {
			errorResponse := dto.InvalidParameterError([]string{"Status is 'public' or 'hidden'"})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Save(&review)

		return c.JSON(
			http.StatusOK,
			dto.ReviewIDResponse{ReviewID: review.ID})
	}
}

func DeleteReviewAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		reviewIDParam := c.Param("reviewId")
		reviewID, err := strconv.Atoi(reviewIDParam)

		if err != nil {
			errorResponse := dto.InvalidParameterError([]string{"ReviewID must be number."})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		var review entity.Review

		if db.Find(&review, reviewID).RecordNotFound() {
			errorResponse := dto.NotFoundError("Review")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		db.Delete(&review, reviewID)

		return c.JSON(
			http.StatusOK,
			dto.ReviewIDResponse{ReviewID: review.ID})
	}
}
