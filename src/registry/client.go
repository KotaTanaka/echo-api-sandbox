package registry

import (
	"gorm.io/gorm"

	clientusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
	clienthandler "github.com/KotaTanaka/echo-api-sandbox/handler/client"
)

type ClientRegistry struct {
	HelloHandler  clienthandler.HelloHandler
	AreaHandler   clienthandler.AreaHandler
	ShopHandler   clienthandler.ShopHandler
	ReviewHandler clienthandler.ReviewHandler
}

func NewClientRegistry(db *gorm.DB) *ClientRegistry {
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

	return &ClientRegistry{
		HelloHandler:  helloHandler,
		AreaHandler:   areaHandler,
		ShopHandler:   shopHandler,
		ReviewHandler: reviewHandler,
	}
}
