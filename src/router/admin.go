package router

import (
	"github.com/KotaTanaka/echo-api-sandbox/registry"
	"github.com/labstack/echo/v4"
)

func AdminRouter(e *echo.Echo, ar *registry.AdminRegistry) {
	// AA-01 エリア登録
	e.POST("/admin/areas", ar.AreaHandler.RegisterArea)
	// AA-02 エリア削除
	e.DELETE("/admin/areas/:areaKey", ar.AreaHandler.DeleteArea)
	// AW-01 Wi-Fiサービス一覧取得・検索
	e.GET("/admin/services", ar.ServiceHandler.GetServiceList)
	// AW-02 Wi-Fiサービス詳細取得
	e.GET("/admin/services/:serviceId", ar.ServiceHandler.GetServiceDetail)
	// AW-03 Wi-Fiサービス登録
	e.POST("/admin/services", ar.ServiceHandler.RegisterService)
	// AW-04 Wi-Fiサービス編集
	e.PUT("/admin/services/:serviceId", ar.ServiceHandler.UpdateService)
	// AW-05 Wi-Fiサービス削除
	e.DELETE("/admin/services/:serviceId", ar.ServiceHandler.DeleteService)
	// AS-01 店舗一覧取得・検索
	e.GET("/admin/shops", ar.ShopHandler.GetShopList)
	// AS-02 店舗詳細取得
	e.GET("/admin/shops/:shopId", ar.ShopHandler.GetShopDetail)
	// AS-03 店舗登録
	e.POST("/admin/shops", ar.ShopHandler.RegisterShop)
	// AS-04 店舗編集
	e.PUT("/admin/shops/:shopId", ar.ShopHandler.UpdateShop)
	// AS-05 店舗削除
	e.DELETE("/admin/shops/:shopId", ar.ShopHandler.DeleteShop)
	// AR-01 レビュー一覧取得・検索
	e.GET("/admin/reviews", ar.ReviewHandler.GetReviewList)
	// AR-02 レビューステータス変更
	e.PUT("/admin/reviews/:reviewId", ar.ReviewHandler.UpdateReviewStatus)
	// AR-03 レビュー削除
	e.DELETE("/admin/reviews/:reviewId", ar.ReviewHandler.DeleteReview)
}
