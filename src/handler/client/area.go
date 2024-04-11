package clienthandler

import (
	"net/http"

	"github.com/labstack/echo"

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
	res, err := h.usecase.GetAreaMaster()
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	return ctx.JSON(http.StatusOK, res)
}
