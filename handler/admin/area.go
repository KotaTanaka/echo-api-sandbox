/*
Package adminhandler 管理API関連ハンドラー
*/
package adminhandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"../../data"
	admindata "../../data/admin"
	"../../model"
)

/*
RegisterAreaAdmin | エリア登録
*/
func RegisterAreaAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindata.RegisterAreaRequestBody)
		errorResponse := new(data.ErrorResponse)

		if err := c.Bind(body); err != nil {
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Invalid Request"
			errorResponse.DetailMessage = []string{err.Error()}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Invalid Parameter"
			errorResponse.DetailMessage = strings.Split(err.(validator.ValidationErrors).Error(), "\n")
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		area := new(model.Area)
		area.AreaKey = body.AreaKey
		area.AreaName = body.AreaName

		db.Create(&area)

		return c.JSON(
			http.StatusOK,
			data.AreaKeyResponse{AreaKey: area.AreaKey})
	}
}
