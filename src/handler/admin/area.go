package adminhandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/KotaTanaka/echo-api-sandbox/model/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/model/dto/admin"
)

func RegisterAreaAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(admindto.RegisterAreaRequest)

		if err := c.Bind(body); err != nil {
			errorResponse := dto.InvalidRequestError([]string{err.Error()})
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if err := c.Validate(body); err != nil {
			errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		area := new(model.Area)
		area.AreaKey = body.AreaKey
		area.AreaName = body.AreaName

		db.Create(&area)

		return c.JSON(
			http.StatusOK,
			dto.AreaKeyResponse{AreaKey: area.AreaKey})
	}
}

func DeleteAreaAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		areaKey := c.Param("areaKey")

		area := model.Area{}
		db.Where("area_key = ?", areaKey).Find(&area)
		db.Delete(&area)

		return c.JSON(
			http.StatusOK,
			dto.AreaKeyResponse{AreaKey: area.AreaKey})
	}
}
