package router

import (
	"category/internal/handler"
	"category/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type RouterHandler struct {
	categoryApi *handler.CategoryHandler
}

func NewRouterHandler(categoryApi *handler.CategoryHandler) *RouterHandler {
	return &RouterHandler{
		categoryApi: categoryApi,
	}
}

func (r *RouterHandler) InitRouter(root *fiber.App, md *middleware.Middleware) {

	// Health check endpoint (registered before middleware to avoid rate limiting)
	root.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "categoryservice",
		})
	})

	root.Use(md.Rl.Handler())
	root.Use(md.Log.Handler())

	categoryGroup := (*root).Group("/category")

	categoryGroup.Get("/get/:id", r.categoryApi.GetCategory)
	categoryGroup.Post("/create", r.categoryApi.CreateCategory)
	categoryGroup.Delete("/delete/:id", r.categoryApi.DeleteCategory)
	categoryGroup.Put("/update/:id", r.categoryApi.UpdateCategory)
	categoryGroup.Get("/list", r.categoryApi.GetCategoryList)
}
