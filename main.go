/*
Find Wi-Fi API main.go
*/
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./data"
	"./handler"
)

/*
connectGorm DBに接続する
*/
func connectGorm() *gorm.DB {
	// TODO 設定ファイルに書く
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "find_wifi_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

/*
Main
*/
func main() {
	e := echo.New()
	db := connectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&data.Service{}, &data.Shop{}, &data.Review{})

	// リクエスト共通処理
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", handler.HelloWorld())
	e.GET("/shops", handler.GetShopsListClient())

	e.Logger.Fatal(e.Start(":1323"))
}
