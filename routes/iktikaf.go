package routes

import (
	"go-fiber-api/database"
	"go-fiber-api/middleware"
	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

func IktikafRoutes(app *fiber.App) {
	api := app.Group("/iktikaf")

	// 📌 CREATE: Bisa diakses tanpa login/token
	api.Post("/", func(c *fiber.Ctx) error {
		var iktikaf models.Iktikaf
		if err := c.BodyParser(&iktikaf); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		database.DB.Create(&iktikaf)
		return c.Status(201).JSON(iktikaf)
	})

	// 📌 READ: Get all registrations (hanya untuk user login)
	api.Get("/", middleware.AuthRequired, func(c *fiber.Ctx) error {
		var iktikaf []models.Iktikaf
		database.DB.Find(&iktikaf)
		return c.JSON(iktikaf)
	})

	// 📌 READ: Get one by ID (hanya untuk user login)
	api.Get("/:id", middleware.AuthRequired, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var iktikaf models.Iktikaf
		if err := database.DB.First(&iktikaf, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Iktikaf not found"})
		}
		return c.JSON(iktikaf)
	})

	// 📌 UPDATE: Hanya bisa diakses oleh user yang login
	api.Put("/:id", middleware.AuthRequired, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var iktikaf models.Iktikaf
		if err := database.DB.First(&iktikaf, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Iktikaf not found"})
		}

		if err := c.BodyParser(&iktikaf); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		database.DB.Save(&iktikaf)
		return c.JSON(iktikaf)
	})

	// 📌 DELETE: Hanya bisa diakses oleh user yang login
	api.Delete("/:id", middleware.AuthRequired, func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := database.DB.Delete(&models.Iktikaf{}, id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete Iktikaf"})
		}
		return c.JSON(fiber.Map{"message": "Iktikaf registration deleted"})
	})
}
