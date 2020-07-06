/*
Package server サーバー全体のセットアップ
*/
package server

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"../handler"
	adminhandler "../handler/admin"
	clienthandler "../handler/client"
)

/*
Router | ルーティング
*/
func Router(e *echo.Echo, db *gorm.DB) {
	// API仕様書の出力
	e.File("/doc", "app/redoc.html")

	// ルーティング
	e.GET("/", handler.Hello())
	e.GET("/areas", clienthandler.GetAreaMasterClient(db))
	e.GET("/shops", clienthandler.GetShopListClient(db))
	e.POST("/admin/areas", adminhandler.RegisterAreaAdmin(db))
	e.DELETE("/admin/areas/:areaKey", adminhandler.DeleteAreaAdmin(db))
	e.GET("/admin/services", adminhandler.GetServiceListAdmin(db))
	e.POST("/admin/services", adminhandler.RegisterServiceAdmin(db))
	e.GET("/admin/shops", adminhandler.GetShopListAdmin(db))
	e.POST("/admin/shops", adminhandler.RegisterShopAdmin(db))
}
