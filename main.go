/*
Find Wi-Fi API main.go
*/
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./handler"
)

func main() {
	e := echo.New()

	// リクエスト共通処理
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", handler.HelloWorld())

	e.Logger.Fatal(e.Start(":1323"))
}
