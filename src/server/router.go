package server

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/KotaTanaka/echo-api-sandbox/handler"
	adminhandler "github.com/KotaTanaka/echo-api-sandbox/handler/admin"
	clienthandler "github.com/KotaTanaka/echo-api-sandbox/handler/client"
)

func Router(e *echo.Echo, db *gorm.DB) {
	// API仕様書の出力
	e.File("/doc", "app/redoc.html")
	// Hello, World!
	e.GET("/", handler.Hello())
	// CA-01 エリアマスタ取得
	e.GET("/areas", clienthandler.GetAreaMasterClient(db))
	// CS-01 エリアに紐づく店舗一覧取得
	e.GET("/shops", clienthandler.GetShopListClient(db))
	// CR-01 店舗に紐づくレビュー一覧取得
	e.GET("/reviews", clienthandler.GetReviewListClient(db))
	// CR-02 店舗へのレビュー投稿
	e.POST("/reviews", clienthandler.CreateReviewClient(db))
	// AA-01 エリア登録
	e.POST("/admin/areas", adminhandler.RegisterAreaAdmin(db))
	// AA-02 エリア削除
	e.DELETE("/admin/areas/:areaKey", adminhandler.DeleteAreaAdmin(db))
	// AW-01 Wi-Fiサービス一覧取得・検索
	e.GET("/admin/services", adminhandler.GetServiceListAdmin(db))
	// AW-02 Wi-Fiサービス詳細取得
	e.GET("/admin/services/:serviceId", adminhandler.GetServiceDetailAdmin(db))
	// AW-03 Wi-Fiサービス登録
	e.POST("/admin/services", adminhandler.RegisterServiceAdmin(db))
	// AW-04 Wi-Fiサービス編集
	e.PUT("/admin/services/:serviceId", adminhandler.UpdateServiceAdmin(db))
	// AW-05 Wi-Fiサービス削除
	e.DELETE("/admin/services/:serviceId", adminhandler.DeleteServiceAdmin(db))
	// AS-01 店舗一覧取得・検索
	e.GET("/admin/shops", adminhandler.GetShopListAdmin(db))
	// AS-02 店舗詳細取得
	e.GET("/admin/shops/:shopId", adminhandler.GetShopDetailAdmin(db))
	// AS-03 店舗登録
	e.POST("/admin/shops", adminhandler.RegisterShopAdmin(db))
	// AS-04 店舗編集
	e.PUT("/admin/shops/:shopId", adminhandler.UpdateShopAdmin(db))
	// AS-05 店舗削除
	e.DELETE("/admin/shops/:shopId", adminhandler.DeleteShopAdmin(db))
	// AR-01 レビュー一覧取得・検索
	e.GET("/admin/reviews", adminhandler.GetReviewListAdmin(db))
	// AR-02 レビューステータス変更
	e.PUT("/admin/reviews/:reviewId", adminhandler.UpdateReviewStatusAdmin(db))
	// AR-03 レビュー削除
	e.DELETE("/admin/reviews/:reviewId", adminhandler.DeleteReviewAdmin(db))
}
