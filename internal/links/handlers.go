package links

import (
	"surl/internal/links/dtos"
	"surl/internal/links/usecase"

	"github.com/labstack/echo/v4"
)

func handleCreate(ctx echo.Context) error {
	dto := new(dtos.CreateLinkDTO)
	if err := ctx.Bind(dto); err != nil {
		return echo.ErrBadRequest
	}
	shortened, err := usecase.CreateShortenedLink(dto.Link)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(201, echo.Map{
		"shortened": shortened, // TODO add base URL
	})
}

func handleFindOne(ctx echo.Context) error {
	return nil
}
