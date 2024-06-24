package db

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/million_dollar_space_programme/exoplanets/configs"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     configs.DATABASE_URL,
		Password: "",
		DB:       0,
	})
}

func GetDBClient(ctx context.Context) (*redis.Client, error) {
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
