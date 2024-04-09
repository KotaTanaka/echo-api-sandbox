package handler

import (
	"net/http"

	"github.com/KotaTanaka/echo-api-sandbox/data"

	"github.com/labstack/echo"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			data.MessageResponse{Message: "Hello, Find Wi-Fi!"})
	}
}
