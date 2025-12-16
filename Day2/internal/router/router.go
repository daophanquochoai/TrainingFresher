package router

import (
	"github.com/go-chi/chi/v5"
	"go-db-demo/internal/handler"
)

func NewRouter(
	userHandler *handler.UserHandler,
	categoryHandler *handler.CategoryHandler,
	productHandler *handler.ProductHandler,
) *chi.Mux {

	r := chi.NewRouter()

	RegisterUserRoutes(r, userHandler)
	RegisterCagoriesRoutes(r, categoryHandler)
	RegistProductRoutes(r, productHandler)
	return r
}
