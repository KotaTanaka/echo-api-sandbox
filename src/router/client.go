package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clienthandler "github.com/KotaTanaka/echo-api-sandbox/handler/client"
)

func ClientRouter(e *echo.Echo, db *gorm.DB) {
	// Hello, World!
	e.GET("/", clienthandler.Hello())
	// CA-01 エリアマスタ取得
	e.GET("/areas", clienthandler.GetAreaMasterClient(db))
	// CS-01 エリアに紐づく店舗一覧取得
	e.GET("/shops", clienthandler.GetShopListClient(db))
	// CR-01 店舗に紐づくレビュー一覧取得
	e.GET("/reviews", clienthandler.GetReviewListClient(db))
	// CR-02 店舗へのレビュー投稿
	e.POST("/reviews", clienthandler.CreateReviewClient(db))
}
