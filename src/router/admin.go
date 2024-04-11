package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	adminhandler "github.com/KotaTanaka/echo-api-sandbox/handler/admin"
)

func AdminRouter(e *echo.Echo, db *gorm.DB) {
	areaHandler := adminhandler.NewAreaHandler(db)
	serviceHandler := adminhandler.NewServiceHandler(db)
	shopHandler := adminhandler.NewShopHandler(db)
	reviewHandler := adminhandler.NewReviewHandler(db)

	// AA-01 エリア登録
	e.POST("/admin/areas", areaHandler.RegisterArea)
	// AA-02 エリア削除
	e.DELETE("/admin/areas/:areaKey", areaHandler.DeleteArea)
	// AW-01 Wi-Fiサービス一覧取得・検索
	e.GET("/admin/services", serviceHandler.GetServiceList)
	// AW-02 Wi-Fiサービス詳細取得
	e.GET("/admin/services/:serviceId", serviceHandler.GetServiceDetail)
	// AW-03 Wi-Fiサービス登録
	e.POST("/admin/services", serviceHandler.RegisterService)
	// AW-04 Wi-Fiサービス編集
	e.PUT("/admin/services/:serviceId", serviceHandler.UpdateService)
	// AW-05 Wi-Fiサービス削除
	e.DELETE("/admin/services/:serviceId", serviceHandler.DeleteService)
	// AS-01 店舗一覧取得・検索
	e.GET("/admin/shops", shopHandler.GetShopList)
	// AS-02 店舗詳細取得
	e.GET("/admin/shops/:shopId", shopHandler.GetShopDetail)
	// AS-03 店舗登録
	e.POST("/admin/shops", shopHandler.RegisterShop)
	// AS-04 店舗編集
	e.PUT("/admin/shops/:shopId", shopHandler.UpdateShop)
	// AS-05 店舗削除
	e.DELETE("/admin/shops/:shopId", shopHandler.DeleteShop)
	// AR-01 レビュー一覧取得・検索
	e.GET("/admin/reviews", reviewHandler.GetReviewList)
	// AR-02 レビューステータス変更
	e.PUT("/admin/reviews/:reviewId", reviewHandler.UpdateReviewStatus)
	// AR-03 レビュー削除
	e.DELETE("/admin/reviews/:reviewId", reviewHandler.DeleteReview)
}
