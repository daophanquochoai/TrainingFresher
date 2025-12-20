package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/handler"
	"go-db-demo/internal/middleware"
)

func RegistProductRoutes(r chi.Router, h *handler.ProductHandler, redis *redis.Client) {
	r.Route("/products", func(r chi.Router) {
		r.Use(middleware.Auth(redis))
		r.Post("/", h.CreateProductHanler)
		r.Get("/{id}", h.GetProductById)
		r.Put("/update/{id}", h.UpdateProductById)
		r.Delete("/delete/{id}", h.DeleteProductById)
	})
}
