package links

import "github.com/labstack/echo/v4"

func DefineRoutes(e *echo.Echo) {
	group := e.Group("/links")
	group.POST("/", handleCreate)
	group.GET("/:slug", handleFindOne)
}
