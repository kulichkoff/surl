package db

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
