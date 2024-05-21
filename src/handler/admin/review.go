package adminhandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
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

func (h *reviewHandler) GetReviewList(ctx echo.Context) error {
	res, errRes := h.usecase.GetReviewList()
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *reviewHandler) UpdateReviewStatus(ctx echo.Context) error {
	reviewIDParam := ctx.Param("reviewId")
	reviewID, err := strconv.Atoi(reviewIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ReviewID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	body := new(admindto.UpdateReviewStatusRequest)
	if err := ctx.Bind(body); err != nil {
		errRes := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	validator.New()
	if err := ctx.Validate(body); err != nil {
		errRes := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.UpdateReviewStatus(reviewID, body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *reviewHandler) DeleteReview(ctx echo.Context) error {
	reviewIDParam := ctx.Param("reviewId")
	reviewID, err := strconv.Atoi(reviewIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ReviewID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.DeleteReview(reviewID)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}
