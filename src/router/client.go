package router

import (
	"github.com/labstack/echo/v4"

	"github.com/KotaTanaka/echo-api-sandbox/registry"
)

func ClientRouter(e *echo.Echo, cr *registry.ClientRegistry) {
	// Hello, World!
	e.GET("/", cr.HelloHandler.Hello)
	// CA-01 エリアマスタ取得
	e.GET("/areas", cr.AreaHandler.GetAreaMaster)
	// CS-01 エリアに紐づく店舗一覧取得
	e.GET("/shops", cr.ShopHandler.GetShopList)
	// CR-01 店舗に紐づくレビュー一覧取得
	e.GET("/reviews", cr.ReviewHandler.GetReviewList)
	// CR-02 店舗へのレビュー投稿
	e.POST("/reviews", cr.ReviewHandler.CreateReview)
}
