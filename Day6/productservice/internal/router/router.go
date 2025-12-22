package router

import (
	"github.com/gofiber/fiber/v2"
	"product/internal/handler"
	"product/internal/middleware"
)

type RouterHandler struct {
	productApi *handler.ProductHandler
}

func NewRouterHandler(productApi *handler.ProductHandler) *RouterHandler {
	return &RouterHandler{productApi: productApi}
}

func (r *RouterHandler) InitRouter(root *fiber.App, md *middleware.Middleware) {

	// Health check endpoint (registered before middleware to avoid rate limiting)
	root.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "productservice",
		})
	})

	root.Use(md.RateLimit.Middleware())
	root.Use(md.Log.Handler())

	productGroup := (*root).Group("/product")

	productGroup.Get("/get/:id", r.productApi.GetProduct)
	productGroup.Get("/list", r.productApi.GetProductByFilter)
	productGroup.Post("/create", r.productApi.CreateProduct)
	productGroup.Put("/update/:id", r.productApi.UpdateProduct)
	productGroup.Delete("/delete/:id", r.productApi.DeleteProduct)
}
