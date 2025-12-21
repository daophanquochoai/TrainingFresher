package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"product/internal/config"
	"time"
)

func NewRedisClient(redisConfig config.Redis) (*redis.Client, error) {
	rd := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel()

	if err := rd.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rd, nil
}
