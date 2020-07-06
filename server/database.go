/*
Package server サーバー全体のセットアップ
*/
package server

import (
	"os"

	"github.com/jinzhu/gorm"

	"../model"
)

/*
ConnectGorm | Gormの接続
*/
func ConnectGorm() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	DBNAME := os.Getenv("MYSQL_DB")
	OPTION := "charset=utf8mb4&loc=Asia%2FTokyo&parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

/*
Migrate | DBの構築
*/
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Area{})
	db.AutoMigrate(&model.Service{})
	db.AutoMigrate(&model.Shop{}).AddForeignKey("service_id", "services(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.Review{}).AddForeignKey("shop_id", "shops(id)", "RESTRICT", "RESTRICT")
}
