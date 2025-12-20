package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/handler"
)

func NewRouter(
	userHandler *handler.UserHandler,
	categoryHandler *handler.CategoryHandler,
	productHandler *handler.ProductHandler,
	redis *redis.Client,
) *chi.Mux {

	r := chi.NewRouter()
	RegisterUserRoutes(r, userHandler, redis)
	RegisterCagoriesRoutes(r, categoryHandler, redis)
	RegistProductRoutes(r, productHandler, redis)
	return r
}
