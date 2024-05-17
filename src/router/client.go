package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	clientusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
	clienthandler "github.com/KotaTanaka/echo-api-sandbox/handler/client"
)

func ClientRouter(e *echo.Echo, db *gorm.DB) {
	areaRepository := repository.NewAreaRepository(db)
	serviceRepository := repository.NewServiceRepository(db)
	shopRepository := repository.NewShopRepository(db)
	reviewRepository := repository.NewReviewRepository(db)

	areaUsecase := clientusecase.NewAreaUsecase(areaRepository)
	shopUsecase := clientusecase.NewShopUsecase(serviceRepository, shopRepository, reviewRepository)
	reviewUsecase := clientusecase.NewReviewUsecase(serviceRepository, shopRepository, reviewRepository)

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
