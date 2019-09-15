/*
Package handler
*/
package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
Hello, World!
*/
func HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
