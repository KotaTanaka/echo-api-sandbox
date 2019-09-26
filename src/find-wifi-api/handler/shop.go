/*
Package handler - Shop 関連
*/
package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"../data"
)

/*
GetShopsListClient - 店舗一覧を取得する
@author kotatanaka
*/
func GetShopsListClient() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := data.ShopListingResponse{
			ShopList: []data.ShopListingResponseElement{
				{
					ShopID:       "6aa3b8b5-1b5c-40c7-85cc-1cde056eb4c2",
					ShopName:     "スターバックス渋谷店",
					WifiName:     "スターバックスWi-Fi",
					Ssid:         []string{},
					Address:      "東京都渋谷区道玄坂1-1-1",
					Acceess:      "JR渋谷駅から徒歩1分",
					Description:  "Wi-Fi・電源完備で、平日の昼間は比較的落ち着いています。",
					ShopType:     "cafe",
					OpeningHours: "9:00-22:00",
					SeatsNum:     100,
					Power:        true,
					ReviewCount:  10,
					Average:      3.5}},
			Total: 50}
		return c.JSON(http.StatusOK, response)
	}
}
