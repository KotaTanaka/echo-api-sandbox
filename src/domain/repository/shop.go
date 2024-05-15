package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/jinzhu/gorm"
)

type ShopRepository interface {
	ListShops() ([]*model.Shop, error)
	ListShopsByServiceID(serviceID int) ([]*model.Shop, error)
	FindShopByID(shopID int) (*model.Shop, error)
	CreateShop(shop *model.Shop) (*model.Shop, error)
	UpdateShop(shop *model.Shop) (*model.Shop, error)
	DeleteShop(shop *model.Shop) error
	CountShopsByServiceID(serviceID int) (*model.Aggregation, error)
}

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepository{db: db}
}

func (r *shopRepository) ListShops() ([]*model.Shop, error) {
	shops := []*model.Shop{}
	r.db.Find(&shops)

	return shops, nil
}

func (r *shopRepository) ListShopsByServiceID(serviceID int) ([]*model.Shop, error) {
	shops := []*model.Shop{}
	r.db.Where("service_id = ?", serviceID).Find(&shops)

	return shops, nil
}

func (r *shopRepository) FindShopByID(shopID int) (*model.Shop, error) {
	var shop model.Shop
	r.db.First(&shop, shopID)

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

func (r *shopRepository) CountShopsByServiceID(serviceID int) (*model.Aggregation, error) {
	var shopCount int
	r.db.Model(&model.Shop{}).Where("service_id = ?", serviceID).Count(&shopCount)

	return &model.Aggregation{
		Count: shopCount,
	}, nil
}
