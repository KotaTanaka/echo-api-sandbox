package clienthandler

import (
	"net/http"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"

	"github.com/labstack/echo"
)

type HelloHandler interface {
	Hello(ctx echo.Context) error
}

type helloHandler struct{}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Hello(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		dto.MessageResponse{
			Message: "Hello, Find Wi-Fi!",
		},
	)
}
