package clienthandler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

func GetAreaMasterClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		areas := []model.Area{}
		db.Find(&areas)

		response := clientdto.AreaMasterResponse{}
		response.AreaList = []clientdto.AreaMasterResponseElement{}

		for _, area := range areas {
			response.AreaList = append(
				response.AreaList, clientdto.AreaMasterResponseElement{
					AreaKey:   area.AreaKey,
					AreaName:  area.AreaName,
					ShopCount: len(area.Shops)})
		}

		return c.JSON(http.StatusOK, response)
	}
}
