package clienthandler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	clientusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/client"
)

type ReviewHandler interface {
	GetReviewList(ctx echo.Context) error
	CreateReview(ctx echo.Context) error
}

type reviewHandler struct {
	usecase clientusecase.ReviewUsecase
}

func NewReviewHandler(usecase clientusecase.ReviewUsecase) ReviewHandler {
	return &reviewHandler{usecase: usecase}
}

func (h *reviewHandler) GetReviewList(ctx echo.Context) error {
	query := &clientdto.ReviewListingQuery{
		ShopID: ctx.QueryParam("shopId"),
	}

	res, err := h.usecase.GetReviewList(query)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *reviewHandler) CreateReview(ctx echo.Context) error {
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

	res, err := h.usecase.CreateReview(body)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
