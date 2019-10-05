/*
Package data - Review関連の構造体
*/
package data

import (
	"github.com/jinzhu/gorm"
)

/*
Review Model
*/
type Review struct {
	gorm.Model
	ShopID        int
	Comment       string `gorm:"size:1000"`
	Evaluation    int
	PuplishStatus bool
}
