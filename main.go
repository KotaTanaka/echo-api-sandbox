/*
Find Wi-Fi API main.go
*/
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/server"
)

/*
Validator | バリデーターの構造体
*/
type Validator struct {
	validator *validator.Validate
}

/*
Validate | バリデーターのセットアップ
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
	db := server.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	server.Migrate(db)

	// ルーティング
	server.Router(e, db)

	// リクエスト共通処理
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":1323"))
}
