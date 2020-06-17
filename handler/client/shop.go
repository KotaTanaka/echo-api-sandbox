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
GetShopListClient | 店舗一覧取得
*/
func GetShopListClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shops := []model.Shop{}
		db.Find(&shops)

		response := clientdata.ShopListingResponse{}
		response.Total = len(shops)

		for _, shop := range shops {
			service := model.Service{}
			db.First(&service, shop.ServiceID)

			response.ShopList = append(
				// TODO SSID: 文字列を配列に変換
				// TODO Average: 評価の平均値の計算
				response.ShopList, clientdata.ShopListingResponseElement{
					ShopID:       shop.ID,
					WifiName:     service.WifiName,
					ServiceLink:  service.Link,
					ShopName:     shop.ShopName,
					Area:         shop.AreaKey,
					Description:  shop.Description,
					Address:      shop.Address,
					Access:       shop.Access,
					SSID:         []string{shop.SSID},
					ShopType:     shop.ShopType,
					OpeningHours: shop.OpeningHours,
					SeatsNum:     shop.SeatsNum,
					HasPower:     shop.HasPower,
					ReviewCount:  len(shop.Reviews),
					Average:      0})
		}

		return c.JSON(http.StatusOK, response)
	}
}
