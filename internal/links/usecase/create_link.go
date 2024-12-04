package usecase

import (
	"context"
	"surl/internal/db"
	"surl/internal/links/entities"
)

func CreateShortenedLink(original string) (string, error) {
	link := entities.NewLink(original)
	// TODO check the same slug doesn't exist
	db.RDB.Set(context.Background(), link.GetShotened(), link.GetOrinial(), 0)
	return link.GetShotened(), nil
}
