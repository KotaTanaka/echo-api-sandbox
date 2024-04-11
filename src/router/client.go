package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	clientusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/client"
	clienthandler "github.com/KotaTanaka/echo-api-sandbox/handler/client"
)

func ClientRouter(e *echo.Echo, db *gorm.DB) {
	areaUsecase := clientusecase.NewAreaUsecase(db)
	shopUsecase := clientusecase.NewShopUsecase(db)
	reviewUsecase := clientusecase.NewReviewUsecase(db)

	helloHandler := clienthandler.NewHelloHandler()
	areaHandler := clienthandler.NewAreaHandler(areaUsecase)
	shopHandler := clienthandler.NewShopHandler(shopUsecase)
	reviewHandler := clienthandler.NewReviewHandler(reviewUsecase)

	// Hello, World!
	e.GET("/", helloHandler.Hello)
	// CA-01 エリアマスタ取得
	e.GET("/areas", areaHandler.GetAreaMaster)
	// CS-01 エリアに紐づく店舗一覧取得
	e.GET("/shops", shopHandler.GetShopList)
	// CR-01 店舗に紐づくレビュー一覧取得
	e.GET("/reviews", reviewHandler.GetReviewList)
	// CR-02 店舗へのレビュー投稿
	e.POST("/reviews", reviewHandler.CreateReview)
}
