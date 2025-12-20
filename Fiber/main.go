package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("UP")
	})

	type CreateUserReq struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	app.Post("/users", func(c *fiber.Ctx) error {
		var req CreateUserReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid json",
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "created",
			"data":    req,
		})
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{"updated_id": id})
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{"deleted_id": id})
	})

	app.Listen(":8080")
}
