package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"product/internal/dto"
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

	ct, cancel := context.WithTimeout(c.Context(), 5*time.Minute)

	defer cancel()

	//create
	created, errCreate := h.s.CreateProduct(ct, &req)

	if errCreate != nil {
		return fiber.NewError(500, errCreate.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(created)
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Id : ", id)
	if id == "" {
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

func (h *ProductHandler) GetProductByFilter(c *fiber.Ctx) error {
	var req dto.PageRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, ErrInvalid)
	}

	// timeout
	ct, cancel := context.WithTimeout(c.Context(), 5*time.Minute)
	defer cancel()

	// get product
	product, err := h.s.GetProductByFilter(ct, &req)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Id : ", id)
	if id == "" {
		return fiber.NewError(400, ErrInvalid)
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(400, ErrInvalid)
	}

	var req model.Product
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, ErrInvalid)
	}

	// timeout
	ct, cancel := context.WithTimeout(c.Context(), 5*time.Minute)
	defer cancel()

	// get product
	product, err := h.s.UpdateProduct(ct, idUUID, &req)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Id : ", id)
	if id == "" {
		return fiber.NewError(400, ErrInvalid)
	}

	// timeout
	ct, cancel := context.WithTimeout(c.Context(), 5*time.Minute)
	defer cancel()

	// get product
	err := h.s.DeleteProduct(ct, id)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Delete product successfully",
	})
}
