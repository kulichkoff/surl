package main

import (
	"surl/internal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.CORS())

	internal.DefineRoutes(app)

	app.HideBanner = true
	app.Logger.Fatal(app.Start(":8080"))
}
