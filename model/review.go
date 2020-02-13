/*
Package model | レビューモデル
*/
package model

import "github.com/jinzhu/gorm"

/*
Review | 店舗レビューモデル
*/
type Review struct {
	gorm.Model
	ShopID        uint
	Comment       string `gorm:"size:1000"`
	Evaluation    int
	PuplishStatus bool
}
