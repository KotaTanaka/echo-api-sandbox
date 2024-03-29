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

	"github.com/KotaTanaka/echo-api-sandbox/data"
	admindata "github.com/KotaTanaka/echo-api-sandbox/data/admin"
	"github.com/KotaTanaka/echo-api-sandbox/model"
)

/*
RegisterAreaAdmin | エリア登録
*/
func RegisterAreaAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindata.RegisterAreaRequestBody)

		if err := c.Bind(body); err != nil {
			errorResponse := data.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := data.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
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

/*
DeleteAreaAdmin | エリア削除
*/
func DeleteAreaAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		areaKey := c.Param("areaKey")

		area := model.Area{}
		db.Where("area_key = ?", areaKey).Find(&area)
		db.Delete(&area)

		return c.JSON(
			http.StatusOK,
			data.AreaKeyResponse{AreaKey: area.AreaKey})
	}
}
