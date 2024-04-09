package clienthandler

import (
	"net/http"

	"github.com/KotaTanaka/echo-api-sandbox/model/dto"

	"github.com/labstack/echo"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			dto.MessageResponse{Message: "Hello, Find Wi-Fi!"})
	}
}
