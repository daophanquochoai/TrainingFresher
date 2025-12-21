package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RequestLoggerMiddleware struct{}

func NewRequestLoggerMiddleware() *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{}
}

func (m *RequestLoggerMiddleware) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Log request info
		duration := time.Since(start)
		fmt.Printf("[%s] %s %s - %d - %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			duration,
		)

		return err
	}
}
