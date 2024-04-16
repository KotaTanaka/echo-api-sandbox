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

type ShopHandler interface {
	GetShopList(ctx echo.Context) error
	GetShopDetail(ctx echo.Context) error
	RegisterShop(ctx echo.Context) error
	UpdateShop(ctx echo.Context) error
	DeleteShop(ctx echo.Context) error
}

type shopHandler struct {
	usecase adminusecase.ShopUsecase
}

func NewShopHandler(usecase adminusecase.ShopUsecase) ShopHandler {
	return &shopHandler{usecase: usecase}
}

func (h *shopHandler) GetShopList(ctx echo.Context) error {
	res, errRes := h.usecase.GetShopList()
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *shopHandler) GetShopDetail(ctx echo.Context) error {
	shopIDParam := ctx.Param("shopId")
	shopID, err := strconv.Atoi(shopIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ShopID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.GetShopDetail(shopID)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *shopHandler) RegisterShop(ctx echo.Context) error {
	body := new(admindto.RegisterShopRequest)
	if err := ctx.Bind(body); err != nil {
		errRes := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	validator.New()
	if err := ctx.Validate(body); err != nil {
		errRes := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.RegisterShop(body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *shopHandler) UpdateShop(ctx echo.Context) error {
	shopIDParam := ctx.Param("shopId")
	shopID, err := strconv.Atoi(shopIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ShopID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	body := new(admindto.UpdateShopRequest)
	if err := ctx.Bind(body); err != nil {
		errRes := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	validator.New()
	if err := ctx.Validate(body); err != nil {
		errRes := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.UpdateShop(shopID, body)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *shopHandler) DeleteShop(ctx echo.Context) error {
	shopIDParam := ctx.Param("shopId")
	shopID, err := strconv.Atoi(shopIDParam)

	if err != nil {
		errRes := dto.InvalidParameterError([]string{"ShopID must be number."})
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	res, errRes := h.usecase.DeleteShop(shopID)
	if errRes != nil {
		return ctx.JSON(errRes.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
