package middleware

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type Middleware struct {
	Rl  *RateLimit
	Log *LoggingMiddleware
}

func NewMiddleware(r *redis.Client) *Middleware {
	return &Middleware{
		Rl:  NewRateLimit(r, 100, 5*time.Minute),
		Log: NewLoggingMiddleware(),
	}
}
