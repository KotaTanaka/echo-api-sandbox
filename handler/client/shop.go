/*
Package clienthandler クライアントAPI関連ハンドラー
*/
package clienthandler

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientdata "../../data/client"
	"../../model"
)

/*
GetShopListClient | エリアに紐付く店舗一覧取得
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

			reviewCount := 0
			db.Model(&model.Review{}).Where("shop_id = ?", shop.ID).Count(&reviewCount)

			response.ShopList = append(
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
					SSID:         strings.Split(shop.SSID, ","),
					ShopType:     shop.ShopType,
					OpeningHours: shop.OpeningHours,
					SeatsNum:     shop.SeatsNum,
					HasPower:     shop.HasPower,
					ReviewCount:  reviewCount,
					Average:      0})
		}

		return c.JSON(http.StatusOK, response)
	}
}
