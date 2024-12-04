package internal

import (
	"surl/internal/links"

	"github.com/labstack/echo/v4"
)

func DefineRoutes(e *echo.Echo) {
	links.DefineRoutes(e)
}
