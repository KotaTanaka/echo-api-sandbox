package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"

	"github.com/KotaTanaka/echo-api-sandbox/infrastructure"
	"github.com/KotaTanaka/echo-api-sandbox/registry"
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
	db, err := infrastructure.ConnectDB()
	if err != nil {
		fmt.Printf("ConnectDB error: %v", err.Error())
		panic(err.Error())
	}
	defer infrastructure.CloseDB(db)

	// DI
	clientRegistry := registry.NewClientRegistry(db)
	adminRegistry := registry.NewAdminRegistry(db)

	// ルーティング
	router.ClientRouter(e, clientRegistry)
	router.AdminRouter(e, adminRegistry)

	// リクエスト共通処理
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":1323"))
}
