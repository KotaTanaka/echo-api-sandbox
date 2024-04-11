package clienthandler

import (
	"net/http"

	"github.com/labstack/echo"

	clientusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/client"
)

type ShopHandler interface {
	GetShopList(ctx echo.Context) error
}

type shopHandler struct {
	usecase clientusecase.ShopUsecase
}

func NewShopHandler(usecase clientusecase.ShopUsecase) ShopHandler {
	return &shopHandler{usecase: usecase}
}

func (h *shopHandler) GetShopList(ctx echo.Context) error {
	res, err := h.usecase.GetShopList()
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
