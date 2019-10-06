/*
Package data - Service関連の構造体
*/
package data

import "github.com/jinzhu/gorm"

/*
Service Model
*/
type Service struct {
	gorm.Model
	WifiName string `gorm:"size:255"`
	Link     string `gorm:"size:255"`
	Shops    []Shop
}
