package usecase

import (
	"context"
	"surl/internal/db"
)

func FindOneLink(slug string) (string, error) {
	link, err := db.RDB.Get(context.Background(), slug).Result()
	return link, err
}
