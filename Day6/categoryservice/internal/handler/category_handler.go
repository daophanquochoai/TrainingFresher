package handler

import (
	"category/internal/dto"
	"category/internal/model"
	"category/internal/service"
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

var ErrInvalid = "Invalid Data"

type CategoryHandler struct {
	s service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{s: s}
}

func (h *CategoryHandler) GetCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(400, ErrInvalid)
	}

	ct, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	category, err := h.s.GetCategory(ct, id)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(category)

}

func (s *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {

	var category model.Category
	if err := ctx.BodyParser(&category); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if category.Name == "" {
		return fiber.NewError(400, ErrInvalid)
	}

	ct, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	cate, err := s.s.CreateCategory(ct, &category)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.Status(201).JSON(cate)
}

func (h *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(400, ErrInvalid)
	}

	var category model.Category
	if err := ctx.BodyParser(&category); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if category.Name == "" {
		fiber.NewError(400, ErrInvalid)
	}

	ct, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	cate, err := h.s.UpdateCategory(ct, id, &category)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(cate)
}

func (h *CategoryHandler) GetCategoryList(ctx *fiber.Ctx) error {

	var pageRequest dto.PageRequest
	if err := ctx.QueryParser(&pageRequest); err != nil {
		return fiber.NewError(400, err.Error())
	}

	ct, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	cate, err := h.s.GetCategoryList(ct, &pageRequest)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(cate)
}

func (h *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(400, ErrInvalid)
	}

	ct, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	err := h.s.DeleteCategory(ct, id)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON("Delete successfully")
}
