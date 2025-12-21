package cache

import (
	"category/internal/config"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(redisConfig config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
