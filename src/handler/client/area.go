package clienthandler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type AreaHandler interface {
	GetAreaMaster(ctx echo.Context) error
}

type areaHandler struct {
	db *gorm.DB
}

func NewAreaHandler(db *gorm.DB) AreaHandler {
	return &areaHandler{db: db}
}

func (ah areaHandler) GetAreaMaster(ctx echo.Context) error {
	areas := []model.Area{}
	ah.db.Find(&areas)

	response := clientdto.AreaMasterResponse{}
	response.AreaList = []clientdto.AreaMasterResponseElement{}

	for _, area := range areas {
		response.AreaList = append(
			response.AreaList,
			clientdto.AreaMasterResponseElement{
				AreaKey:   area.AreaKey,
				AreaName:  area.AreaName,
				ShopCount: len(area.Shops),
			},
		)
	}

	return ctx.JSON(http.StatusOK, response)
}
