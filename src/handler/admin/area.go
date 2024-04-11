package adminhandler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
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
	validator.New()
	body := new(admindto.RegisterAreaRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	res, err := h.areaUsecase.RegisterArea(body)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *areaHandler) DeleteArea(ctx echo.Context) error {
	query := &admindto.DeleteAreaQuery{
		AreaKey: ctx.Param("areaKey"),
	}

	res, err := h.areaUsecase.DeleteArea(query)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
