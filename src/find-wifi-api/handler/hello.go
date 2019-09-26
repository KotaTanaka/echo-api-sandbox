/*
Package handler - Hello,World!
*/
package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
HelloWorld - Hello,World!を表示する
@author kotatanaka
*/
func HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
