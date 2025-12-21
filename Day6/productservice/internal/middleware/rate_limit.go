package middleware

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRateLimit struct {
	r          *redis.Client
	max        int
	expiration time.Duration
}

func NewRateLimit(r *redis.Client, max int, expiration time.Duration) *RedisRateLimit {
	return &RedisRateLimit{
		r:          r,
		max:        max,
		expiration: expiration,
	}
}

func (rl *RedisRateLimit) Middleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		c := context.Background()

		key := fmt.Sprintf("rate_limit:%s", ctx.IP())

		count, err := rl.r.Get(c, key).Int()
		if err != nil && err != redis.Nil {
			return ctx.Next()
		}

		if count >= rl.max {
			// Get TTL để biết khi nào reset
			ttl, _ := rl.r.TTL(c, key).Result()

			ctx.Set("X-RateLimit-Limit", fmt.Sprintf("%d", rl.max))
			ctx.Set("X-RateLimit-Remaining", "0")
			ctx.Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(ttl).Unix()))

			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "Rate limit exceeded",
				"retry_after": int(ttl.Seconds()),
			})
		}

		pipe := rl.r.Pipeline()
		pipe.Incr(c, key)

		if count == 0 {
			pipe.Expire(c, key, rl.expiration)
		}

		_, err = pipe.Exec(c)
		if err != nil {
			// Nếu Redis lỗi, cho phép request đi qua
			return ctx.Next()
		}

		remaining := rl.max - count - 1
		if remaining < 0 {
			remaining = 0
		}

		ctx.Set("X-RateLimit-Limit", fmt.Sprintf("%d", rl.max))
		ctx.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		ctx.Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(rl.expiration).Unix()))

		return ctx.Next()
	}
}
