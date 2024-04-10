package model

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model
	ShopID        uint
	Comment       string `gorm:"size:1000"`
	Evaluation    int
	PublishStatus bool
}
