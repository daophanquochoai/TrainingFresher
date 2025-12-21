package middleware

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type Middleware struct {
	RateLimit *RedisRateLimit
	Log       *RequestLoggerMiddleware
}

func NewMiddleware(rl *redis.Client) *Middleware {
	return &Middleware{
		RateLimit: NewRateLimit(rl, 100, 1*time.Minute),
		Log:       NewRequestLoggerMiddleware(),
	}
}
