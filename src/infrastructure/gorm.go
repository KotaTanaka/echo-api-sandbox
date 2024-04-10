package infrastructure

import (
	"fmt"
	"os"

	"github.com/KotaTanaka/echo-api-sandbox/model/entity"
	"github.com/jinzhu/gorm"
)

func ConnectGorm() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	DB_NAME := os.Getenv("MYSQL_DB")
	OPTION := "charset=utf8mb4&loc=Asia%2FTokyo&parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Printf("ConnectGorm error: %v", err.Error())
		panic(err.Error())
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.Area{})
	db.AutoMigrate(&entity.Service{})
	db.AutoMigrate(&entity.Shop{}).AddForeignKey("service_id", "services(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&entity.Review{}).AddForeignKey("shop_id", "shops(id)", "RESTRICT", "RESTRICT")
}
