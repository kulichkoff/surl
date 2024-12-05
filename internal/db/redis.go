package db

import (
	"fmt"
	"surl/internal/config"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func init() {
	cfg := config.GetConfig()
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.GetRedisHost(), cfg.GetRedisPort()),
		Password: "",
		DB:       0,
	})
}
