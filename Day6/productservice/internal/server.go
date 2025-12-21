package internal

//
//import (
//	"encoding/json"
//	"github.com/gofiber/fiber/v2"
//	"github.com/gofiber/fiber/v2/middleware/logger"
//	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
//	"github.com/google/wire"
//	"product/internal/cache"
//	"product/internal/config"
//	"product/internal/db"
//	"product/internal/handler"
//	"product/internal/middleware"
//	"product/internal/repository"
//	"product/internal/router"
//	"product/internal/service"
//	"time"
//)
//
//func New() (*fiber.App, error) {
//	panic(wire.Build(
//		config.Set,
//		cache.Set,
//		db.Set,
//		repository.Set,
//		middleware.Set,
//		service.Set,
//		handler.Set,
//		router.Set,
//		NewServer,
//	))
//}
//
//func NewServer(router *router.RouterHandler, md *middleware.Middleware) *fiber.App {
//	app := fiber.New(fiber.Config{
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 5 * time.Second,
//		JSONDecoder:  json.Unmarshal,
//		JSONEncoder:  json.Marshal,
//	})
//
//	app.Use(logger.New())
//
//	recoveryConfig := recoverFiber.ConfigDefault
//	app.Use(recoverFiber.New(recoveryConfig))
//
//	router.InitRouter(app, md)
//
//	return app
//}
