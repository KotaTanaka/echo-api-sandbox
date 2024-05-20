package adminhandler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	adminusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/admin"
)

type AreaHandler interface {
	RegisterArea(ctx echo.Context) error
	DeleteArea(ctx echo.Context) error
}

type areaHandler struct {
	areaUsecase adminusecase.AreaUsecase
}

func NewAreaHandler(areaUsecase adminusecase.AreaUsecase) AreaHandler {
	return &areaHandler{areaUsecase: areaUsecase}
}

func (h *areaHandler) RegisterArea(ctx echo.Context) error {
	body := new(admindto.RegisterAreaRequest)
	if err := ctx.Bind(body); err != nil {
		errRes := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	validator.New()
	if err := ctx.Validate(body); err != nil {
		errRes := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.areaUsecase.RegisterArea(body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *areaHandler) DeleteArea(ctx echo.Context) error {
	query := &admindto.DeleteAreaQuery{
		AreaKey: ctx.Param("areaKey"),
	}

	res, errRes := h.areaUsecase.DeleteArea(query)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}
