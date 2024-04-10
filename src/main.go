package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/infrastructure"
	"github.com/KotaTanaka/echo-api-sandbox/router"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	e := echo.New()

	// バリデーター初期化
	e.Validator = &Validator{validator: validator.New()}

	// DB接続
	db := infrastructure.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	infrastructure.MigrateDB(db)

	// ルーティング
	router.ClientRouter(e, db)
	router.AdminRouter(e, db)

	// リクエスト共通処理
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":1323"))
}
