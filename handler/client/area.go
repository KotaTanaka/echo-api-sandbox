/*
Package clienthandler クライアントAPI関連ハンドラー
*/
package clienthandler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientdata "../../data/client"
	"../../model"
)

/*
GetAreaMasterClient | エリアマスタ取得
*/
func GetAreaMasterClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		areas := []model.Area{}
		db.Find(&areas)

		response := clientdata.AreaMasterResponse{}
		response.AreaList = []clientdata.AreaMasterResponseElement{}

		for _, area := range areas {
			response.AreaList = append(
				response.AreaList, clientdata.AreaMasterResponseElement{
					AreaKey:   area.AreaKey,
					AreaName:  area.AreaName,
					ShopCount: len(area.Shops)})
		}

		return c.JSON(http.StatusOK, response)
	}
}
