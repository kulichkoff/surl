package usecase

import (
	"context"
	"surl/internal/db"
	"surl/internal/links/entities"
)

func CreateShortenedLink(original string) (string, error) {
	link := entities.NewLink(original)
	// TODO check the same slug doesn't exist
	if err := db.RDB.Set(context.Background(), link.GetShotened(), link.GetOrinial(), 0).Err(); err != nil {
		return "", err
	}
	return link.GetShotened(), nil
}
