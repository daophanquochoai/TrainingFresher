package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/handler"
	"go-db-demo/internal/middleware"
)

func RegisterUserRoutes(r chi.Router, h *handler.UserHandler, redis *redis.Client) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Auth(redis))
		r.Post("/", h.CreateUser)
		r.Get("/get/{id}", h.GetUserByID)
		r.Delete("/delete/{id}", h.DeleteUserByID)
		r.Put("/update/{id}", h.UpdateUserByID)
	})
}
