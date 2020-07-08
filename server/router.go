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
	// Hello, World!
	e.GET("/", handler.Hello())
	// CA-01 エリアマスタ取得
	e.GET("/areas", clienthandler.GetAreaMasterClient(db))
	// CS-01 エリアに紐づく店舗一覧取得
	e.GET("/shops", clienthandler.GetShopListClient(db))
	// CR-02 店舗へのレビュー投稿
	e.POST("/reviews", clienthandler.CreateReviewClient(db))
	// AA-01 エリア登録
	e.POST("/admin/areas", adminhandler.RegisterAreaAdmin(db))
	// AA-02 エリア削除
	e.DELETE("/admin/areas/:areaKey", adminhandler.DeleteAreaAdmin(db))
	// AW-01 Wi-Fiサービス一覧取得・検索
	e.GET("/admin/services", adminhandler.GetServiceListAdmin(db))
	// AW-03 Wi-Fiサービス登録
	e.POST("/admin/services", adminhandler.RegisterServiceAdmin(db))
	// AS-01 店舗一覧取得・検索
	e.GET("/admin/shops", adminhandler.GetShopListAdmin(db))
	// AS-03 店舗登録
	e.POST("/admin/shops", adminhandler.RegisterShopAdmin(db))
}
