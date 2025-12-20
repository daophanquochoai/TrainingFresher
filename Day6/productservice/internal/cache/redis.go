package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisConfig struct {
	host     string
	password string
	db       int
}

func NewRedisClient(redisConfig RedisConfig) (*redis.Client, error) {
	rd := redis.NewClient(&redis.Options{
		Addr:     redisConfig.host,
		Password: redisConfig.password,
		DB:       redisConfig.db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel()

	if err := rd.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rd, nil
}
