package router

import (
	"github.com/go-chi/chi/v5"
	"go-db-demo/internal/handler"
)

func RegisterUserRoutes(r chi.Router, h *handler.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/{id}", h.GetUserByID)
	})
}
