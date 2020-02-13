/*
Package model | サービスモデル
*/
package model

import "github.com/jinzhu/gorm"

/*
Service | Wi-Fiサービスモデル
*/
type Service struct {
	gorm.Model
	WifiName string `gorm:"size:255"`
	Link     string `gorm:"size:255"`
	Shops    []Shop
}
