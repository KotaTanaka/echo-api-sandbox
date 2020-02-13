/*
Package data | 店舗レビュー関連の構造体
*/
package data

import (
	"github.com/jinzhu/gorm"
)

/*
Review | Model 店舗レビューテーブル
*/
type Review struct {
	gorm.Model
	ShopID        uint
	Comment       string `gorm:"size:1000"`
	Evaluation    int
	PuplishStatus bool
}

/*
ReviewIDResponse | レビューIDのみのレスポンス
*/
type ReviewIDResponse struct {
	ReviewID uint `json:"reviewId"`
}
