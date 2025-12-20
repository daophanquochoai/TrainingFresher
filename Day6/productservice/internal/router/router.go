package router

import (
	"github.com/gofiber/fiber/v2"
	"product/internal/handler"
)

type RouterHandler struct {
	productApi *handler.ProductHandler
}

func NewRouterHandler(productApi *handler.ProductHandler) *RouterHandler {
	return &RouterHandler{productApi: productApi}
}

func (r RouterHandler) InitRouter(root *fiber.App) {
	productGroup := (*root).Group("/product")

	productGroup.Get("/list", r.productApi.GetProduct)
}
