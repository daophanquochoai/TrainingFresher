package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"product/internal/model"
	"product/internal/service"
	"time"
)

var ErrInvalid string = "Invalid Body"

type ProductHandler struct {
	s service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{
		s: s,
	}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req model.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, ErrInvalid)
	}
	// mapping
	product := model.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	ct, cancel := context.WithTimeout(c.Context(), 5*time.Minute)

	defer cancel()

	//create
	created, errCreate := h.s.CreateProduct(ct, &product)

	if errCreate != nil {
		return fiber.NewError(500, errCreate.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(created)
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id != "" {
		return fiber.NewError(400, ErrInvalid)
	}

	// timeout
	ct, cancel := context.WithTimeout(c.Context(), 5*time.Minute)
	defer cancel()

	// get product
	product, err := h.s.GetProduct(ct, id)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(product)
}
