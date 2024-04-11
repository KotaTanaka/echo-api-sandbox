package clienthandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type ReviewHandler interface {
	GetReviewList(ctx echo.Context) error
	CreateReview(ctx echo.Context) error
}

type reviewHandler struct {
	db *gorm.DB
}

func NewReviewHandler(db *gorm.DB) ReviewHandler {
	return &reviewHandler{db: db}
}

func (rh reviewHandler) GetReviewList(ctx echo.Context) error {
	shopID := ctx.QueryParam("shopId")

	var shop model.Shop
	var service model.Service
	var reviews []model.Review

	if rh.db.Find(&shop, shopID).Related(&reviews).RecordNotFound() {
		errorResponse := dto.NotFoundError("Shop")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	rh.db.First(&service, shop.ID)

	response := clientdto.ReviewListingResponse{}
	response.ShopID = shop.ID
	response.ShopName = shop.ShopName
	response.ServiceID = service.ID
	response.WifiName = service.WifiName
	response.Total = len(reviews)
	response.ReviewList = []clientdto.ReviewListingResponseElement{}

	var evaluationSum int
	for _, review := range reviews {
		evaluationSum += review.Evaluation
		response.ReviewList = append(
			response.ReviewList,
			clientdto.ReviewListingResponseElement{
				ReviewID:   review.ID,
				Comment:    review.Comment,
				Evaluation: review.Evaluation,
				Status:     review.PublishStatus,
				CreatedAt:  review.CreatedAt,
			},
		)
	}

	if response.Total > 0 {
		response.Average = float32(evaluationSum) / float32(response.Total)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (rh reviewHandler) CreateReview(ctx echo.Context) error {
	validator.New()
	body := new(clientdto.CreateReviewRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	shop := model.Shop{}
	rh.db.First(&shop, body.ShopID)

	review := new(model.Review)
	review.ShopID = shop.ID
	review.Comment = body.Comment
	review.Evaluation = body.Evaluation

	rh.db.Create(&review)

	return ctx.JSON(
		http.StatusOK,
		dto.ReviewIDResponse{
			ReviewID: review.ID,
		},
	)
}
