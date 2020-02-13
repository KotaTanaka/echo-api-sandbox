/*
Package handler | Hello, Find Wi-Fi!
*/
package handler

import (
	"net/http"

	"../data"
	"github.com/labstack/echo"
)

/*
Hello | Hello, Find Wi-Fi! 文字列を表示する
*/
func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			data.MessageResponse{Message: "Hello, Find Wi-Fi!"})
	}
}
