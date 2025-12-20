package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/handler"
	"go-db-demo/internal/middleware"
)

func RegisterCagoriesRoutes(r chi.Router, h *handler.CategoryHandler, redis *redis.Client) {
	r.Route("/categories", func(r chi.Router) {
		r.Use(middleware.Auth(redis))
		r.Post("/", h.CreateCategory)
		r.Get("/{id}", h.GetCategoryById)
		r.Put("/update/{id}", h.UpdateCategoryById)
		r.Delete("/delete/{id}", h.DeleteCategoryById)
		r.Put("/products/add", h.AddProductIntoCategory)
	})
}
