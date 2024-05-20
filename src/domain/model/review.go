package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ShopID        uint
	Shop          Shop
	Comment       string `gorm:"size:1000"`
	Evaluation    int
	PublishStatus bool
}
