package adminhandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type AreaHandler interface {
	RegisterArea(ctx echo.Context) error
	DeleteArea(ctx echo.Context) error
}

type areaHandler struct {
	db *gorm.DB
}

func NewAreaHandler(db *gorm.DB) AreaHandler {
	return &areaHandler{db: db}
}

func (ah areaHandler) RegisterArea(ctx echo.Context) error {
	validator.New()
	body := new(admindto.RegisterAreaRequest)

	if err := ctx.Bind(body); err != nil {
		errorResponse := dto.InvalidRequestError([]string{err.Error()})
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := ctx.Validate(body); err != nil {
		errorResponse := dto.InvalidParameterError(strings.Split(err.(validator.ValidationErrors).Error(), "\n"))
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	area := new(model.Area)
	area.AreaKey = body.AreaKey
	area.AreaName = body.AreaName

	ah.db.Create(&area)

	return ctx.JSON(
		http.StatusOK,
		dto.AreaKeyResponse{
			AreaKey: area.AreaKey,
		},
	)
}

func (ah areaHandler) DeleteArea(ctx echo.Context) error {
	areaKey := ctx.Param("areaKey")

	area := model.Area{}
	ah.db.Where("area_key = ?", areaKey).Find(&area)
	ah.db.Delete(&area)

	return ctx.JSON(
		http.StatusOK,
		dto.AreaKeyResponse{
			AreaKey: area.AreaKey,
		},
	)
}
