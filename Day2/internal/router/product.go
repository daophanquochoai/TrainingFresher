package router

import (
	"github.com/go-chi/chi/v5"
	"go-db-demo/internal/handler"
)

func RegistProductRoutes(r chi.Router, h *handler.ProductHandler) {
	r.Route("/products", func(r chi.Router) {
		r.Post("/", h.CreateProductHanler)
		r.Get("/{id}", h.GetProductById)
	})
}
