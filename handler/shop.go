/*
Package handler --- Wi-Fi提供店舗関連ハンドラー
*/
package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"../data"
)

/*
GetShopsListClient --- 店舗一覧取得
@author kotatanaka
*/
func GetShopsListClient(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		shops := []data.Shop{}
		allShops := db.Find(&shops)
		// response := data.ShopListingResponse{
		// 	ShopList: []data.ShopListingResponseElement{
		// 		{
		// 			ShopID:       "6aa3b8b5-1b5c-40c7-85cc-1cde056eb4c2",
		// 			ShopName:     "スターバックス渋谷店",
		// 			WifiName:     "スターバックスWi-Fi",
		// 			Ssid:         []string{},
		// 			Address:      "東京都渋谷区道玄坂1-1-1",
		// 			Acceess:      "JR渋谷駅から徒歩1分",
		// 			Description:  "Wi-Fi・電源完備で、平日の昼間は比較的落ち着いています。",
		// 			ShopType:     "cafe",
		// 			OpeningHours: "9:00-22:00",
		// 			SeatsNum:     100,
		// 			Power:        true,
		// 			ReviewCount:  10,
		// 			Average:      3.5}},
		// 	Total: 50}
		return c.JSON(http.StatusOK, allShops)
	}
}

/*
RegisterShopAdmin --- 店舗登録
@author kotatanaka
*/
func RegisterShopAdmin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		validator.New()
		body := new(data.RegisterShopRequestBody)
		message := new(data.MessageResponse)

		if err := c.Bind(body); err != nil {
			message.Message = err.Error()
			return c.JSON(http.StatusBadRequest, message)
		}

		if err := c.Validate(body); err != nil {
			message.Message = err.(validator.ValidationErrors).Error()
			return c.JSON(http.StatusBadRequest, message)
		}

		shop := new(data.Shop)
		shop.ServiceID = body.ServiceID
		shop.SSID = body.SSID
		shop.ShopName = body.ShopName
		shop.Description = body.Description
		shop.Address = body.Address
		shop.ShopType = body.ShopType
		shop.OpeningHours = body.OpeningHours
		shop.SeatsNum = body.SeatsNum
		shop.HasPower = body.HasPower

		db.Create(&shop)

		return c.JSON(
			http.StatusOK,
			data.ShopIDResponse{ShopID: shop.ID})
	}
}
