package internal

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
	"product/internal/router"
	"time"
)

func NewServer(
	router *router.RouterHandler,
) {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(logger.New())

	recoveryConfig := recoverFiber.ConfigDefault
	app.Use(recoverFiber.New(recoveryConfig))

	router.InitRouter(&app)
}
