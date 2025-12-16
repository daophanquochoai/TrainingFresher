package router

import (
	"github.com/go-chi/chi/v5"
	"go-db-demo/internal/handler"
)

func RegisterCagoriesRoutes(r chi.Router, h *handler.CategoryHandler) {
	r.Route("/categories", func(r chi.Router) {
		r.Post("/", h.CreateCategory)
		r.Get("/{id}", h.GetCategoryById)
	})
}
