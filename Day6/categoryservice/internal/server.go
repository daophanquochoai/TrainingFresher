package internal

//
//import (
//	"category/internal/cache"
//	"category/internal/config"
//	"category/internal/db"
//	"category/internal/grpc/client"
//	"category/internal/grpc/provider"
//	"category/internal/handler"
//	"category/internal/middleware"
//	"category/internal/repository"
//	"category/internal/router"
//	"category/internal/service"
//	"encoding/json"
//	"github.com/gofiber/fiber/v2"
//	"github.com/gofiber/fiber/v2/middleware/logger"
//	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
//	"github.com/google/wire"
//	"time"
//)
//
//func New() (*fiber.App, func(), error) {
//	panic(wire.Build(
//		config.Set,
//		cache.SetCache,
//		db.Set,
//		handler.Set,
//		provider.Set,
//		client.Set,
//		middleware.Set,
//		repository.Set,
//		router.Set,
//		service.Set,
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
