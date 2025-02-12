package main

import (
	"Encode/config"
	uControll "Encode/controller/user"
	"Encode/model/user"
	"Encode/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	m := user.MyClientModel{Connection: db}
	c := uControll.ClientController{Model: m}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, c)

	e.Logger.Fatal(e.Start(":8080"))
}
