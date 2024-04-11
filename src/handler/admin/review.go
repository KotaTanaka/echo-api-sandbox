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

type ReviewHandler interface {
	GetReviewList(ctx echo.Context) error
	UpdateReviewStatus(ctx echo.Context) error
	DeleteReview(ctx echo.Context) error
}

type reviewHandler struct {
	db *gorm.DB
}

func NewReviewHandler(db *gorm.DB) ReviewHandler {
	return &reviewHandler{db: db}
}

func (rh reviewHandler) GetReviewList(ctx echo.Context) error {
	reviews := []model.Review{}
	rh.db.Find(&reviews)

	response := admindto.ReviewListingResponse{}
	response.Total = len(reviews)
	response.ReviewList = []admindto.ReviewListingResponseElement{}

	for _, review := range reviews {
		shop := model.Shop{}
		rh.db.First(&shop, review.ShopID)

		service := model.Service{}
		rh.db.First(&service, shop.ID)

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
				DeletedAt:  review.DeletedAt,
			},
		)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (rh reviewHandler) UpdateReviewStatus(ctx echo.Context) error {
	validator.New()

	reviewIDParam := ctx.Param("reviewId")
	reviewID, err := strconv.Atoi(reviewIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ReviewID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var review model.Review

	if rh.db.Find(&review, reviewID).RecordNotFound() {
		errorResponse := dto.NotFoundError("Review")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	body := new(admindto.UpdateReviewStatusRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if body.Status == "public" {
		review.PublishStatus = true
	} else if body.Status == "hidden" {
		review.PublishStatus = false
	} else {
		errorResponse := dto.InvalidParameterError([]string{"Status is 'public' or 'hidden'"})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	rh.db.Save(&review)

	return ctx.JSON(
		http.StatusOK,
		dto.ReviewIDResponse{
			ReviewID: review.ID,
		},
	)
}

func (rh reviewHandler) DeleteReview(ctx echo.Context) error {
	reviewIDParam := ctx.Param("reviewId")
	reviewID, err := strconv.Atoi(reviewIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ReviewID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	var review model.Review

	if rh.db.Find(&review, reviewID).RecordNotFound() {
		errorResponse := dto.NotFoundError("Review")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	rh.db.Delete(&review, reviewID)

	return ctx.JSON(
		http.StatusOK,
		dto.ReviewIDResponse{
			ReviewID: review.ID,
		},
	)
}
