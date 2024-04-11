package adminhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	adminusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/admin"
)

type ReviewHandler interface {
	GetReviewList(ctx echo.Context) error
	UpdateReviewStatus(ctx echo.Context) error
	DeleteReview(ctx echo.Context) error
}

type reviewHandler struct {
	usecase adminusecase.ReviewUsecase
}

func NewReviewHandler(usecase adminusecase.ReviewUsecase) ReviewHandler {
	return &reviewHandler{usecase: usecase}
}

func (h reviewHandler) GetReviewList(ctx echo.Context) error {
	res, err := h.usecase.GetReviewList()
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h reviewHandler) UpdateReviewStatus(ctx echo.Context) error {
	validator.New()

	reviewIDParam := ctx.Param("reviewId")
	reviewID, err := strconv.Atoi(reviewIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ReviewID must be number."})
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

	res, errRes := h.usecase.UpdateReviewStatus(reviewID, body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h reviewHandler) DeleteReview(ctx echo.Context) error {
	reviewIDParam := ctx.Param("reviewId")
	reviewID, err := strconv.Atoi(reviewIDParam)

	if err != nil {
		errorResponse := dto.InvalidParameterError([]string{"ReviewID must be number."})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	res, errRes := h.usecase.DeleteReview(reviewID)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
