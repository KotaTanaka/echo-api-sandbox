package registry

import (
	"gorm.io/gorm"

	adminusecase "github.com/KotaTanaka/echo-api-sandbox/application/usecase/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
	adminhandler "github.com/KotaTanaka/echo-api-sandbox/handler/admin"
)

type AdminRegistry struct {
	AreaHandler    adminhandler.AreaHandler
	ServiceHandler adminhandler.ServiceHandler
	ShopHandler    adminhandler.ShopHandler
	ReviewHandler  adminhandler.ReviewHandler
}

func NewAdminRegistry(db *gorm.DB) *AdminRegistry {
	areaRepository := repository.NewAreaRepository(db)
	serviceRepository := repository.NewServiceRepository(db)
	shopRepository := repository.NewShopRepository(db)
	reviewRepository := repository.NewReviewRepository(db)

	areaUsecase := adminusecase.NewAreaUsecase(areaRepository)
	serviceUsecase := adminusecase.NewServiceUsecase(serviceRepository, shopRepository, reviewRepository)
	shopUsecase := adminusecase.NewShopUsecase(serviceRepository, shopRepository, reviewRepository)
	reviewUsecase := adminusecase.NewReviewUsecase(serviceRepository, shopRepository, reviewRepository)

	areaHandler := adminhandler.NewAreaHandler(areaUsecase)
	serviceHandler := adminhandler.NewServiceHandler(serviceUsecase)
	shopHandler := adminhandler.NewShopHandler(shopUsecase)
	reviewHandler := adminhandler.NewReviewHandler(reviewUsecase)

	return &AdminRegistry{
		AreaHandler:    areaHandler,
		ServiceHandler: serviceHandler,
		ShopHandler:    shopHandler,
		ReviewHandler:  reviewHandler,
	}
}
