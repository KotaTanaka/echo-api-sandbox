package clienthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	clientusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/client"
)

type AreaHandler interface {
	GetAreaMaster(ctx echo.Context) error
}

type areaHandler struct {
	usecase clientusecase.AreaUsecase
}

func NewAreaHandler(usecase clientusecase.AreaUsecase) AreaHandler {
	return &areaHandler{usecase: usecase}
}

func (h *areaHandler) GetAreaMaster(ctx echo.Context) error {
	res, errRes := h.usecase.GetAreaMaster()
	if errRes != nil {
		return ctx.JSON(errRes.Code, errRes)
	}

	return ctx.JSON(http.StatusOK, res)
}
