package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(host string, password string, db int) (*redis.Client, error) {
	redis := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := redis.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return redis, nil
}
