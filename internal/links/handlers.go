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
		ctx.Logger().Errorf("Failed to create shortened link: %v", err)
		return echo.ErrInternalServerError
	}

	return ctx.JSON(201, echo.Map{
		"shortened": shortened, // TODO add base URL
	})
}

func handleFindOne(ctx echo.Context) error {
	slug := ctx.Param("slug")
	link, err := usecase.FindOneLink(slug)
	if err != nil {
		ctx.Logger().Errorf("Find one link error: %v", err)
		ctx.Logger().Printf("slug: %s; link: %s", slug, link)
		return echo.ErrInternalServerError
	}
	if link == "" {
		return echo.ErrNotFound
	}
	return ctx.JSON(200, echo.Map{
		"link": link,
	})
}
