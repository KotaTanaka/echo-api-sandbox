/*
Find Wi-Fi API main.go
*/
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	"./data"
	"./handler"
)

/*
ConnectGorm --- DBのセットアップ
*/
func ConnectGorm() *gorm.DB {
	// TODO 設定ファイルに書く
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "find_wifi_db"
	OPTION := "parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

/*
Validator --- バリデーターのセットアップのためのラッピング
*/
type Validator struct {
	validator *validator.Validate
}

/*
Validate --- バリデーターのセットアップ
*/
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

/*
Main
*/
func main() {
	e := echo.New()

	// バリデーターのセットアップ
	e.Validator = &Validator{validator: validator.New()}

	// DBのセットアップ
	db := ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB")
	db.AutoMigrate(&data.Service{})
	db.AutoMigrate(&data.Shop{}).AddForeignKey("service_id", "services(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&data.Review{}).AddForeignKey("shop_id", "shops(id)", "RESTRICT", "RESTRICT")

	// リクエスト共通処理
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", handler.HelloWorld())
	e.GET("/shops", handler.GetShopListClient(db))
	e.GET("/admin/services", handler.GetServiceListAdmin(db))
	e.POST("/admin/services", handler.RegisterServiceAdmin(db))
	e.POST("/admin/shops", handler.RegisterShopAdmin(db))

	e.Logger.Fatal(e.Start(":1323"))
}
