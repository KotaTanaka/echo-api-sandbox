package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"gorm.io/gorm"
)

type ShopRepository interface {
	ListShops() ([]*model.Shop, error)
	FindShopByID(shopID int) (*model.Shop, error)
	CreateShop(shop *model.Shop) (*model.Shop, error)
	UpdateShop(shop *model.Shop) (*model.Shop, error)
	DeleteShop(shop *model.Shop) error
}

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepository{db: db}
}

func (r *shopRepository) ListShops() ([]*model.Shop, error) {
	shops := []*model.Shop{}
	r.db.Preload("Area").Preload("Service").Find(&shops)

	return shops, nil
}

func (r *shopRepository) FindShopByID(shopID int) (*model.Shop, error) {
	var shop model.Shop
	r.db.Preload("Area").Preload("Service").Preload("Reviews").First(&shop, shopID)

	return &shop, nil
}

func (r *shopRepository) CreateShop(shop *model.Shop) (*model.Shop, error) {
	r.db.Create(&shop)

	return shop, nil
}

func (r *shopRepository) UpdateShop(shop *model.Shop) (*model.Shop, error) {
	r.db.Save(&shop)

	return shop, nil
}

func (r *shopRepository) DeleteShop(shop *model.Shop) error {
	r.db.Delete(&shop)

	return nil
}
