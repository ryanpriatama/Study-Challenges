package main

import (
	"SC/auth"
	"SC/config"
	"SC/route"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init_DB()
	config.InitPort()
	auth.LogMiddlewares((e))
	route.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
