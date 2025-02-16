package routes

import (
	"go-fiber-api/database"
	"go-fiber-api/middleware"
	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	api := app.Group("/books")

	// Get all books
	api.Get("/", middleware.AuthRequired, func(c *fiber.Ctx) error {
		var books []models.Book
		database.DB.Find(&books)
		return c.JSON(books)
	})

	// Get single book by ID
	api.Get("/:id", middleware.AuthRequired, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.Book
		if err := database.DB.First(&book, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.JSON(book)
	})

	// Create a new book
	api.Post("/", middleware.AuthRequired, func(c *fiber.Ctx) error {
		var book models.Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		database.DB.Create(&book)
		return c.Status(201).JSON(book)
	})

	// Update a book by ID
	api.Put("/:id", middleware.AuthRequired, func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.Book
		if err := database.DB.First(&book, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}

		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		database.DB.Save(&book)
		return c.JSON(book)
	})

	// Delete a book by ID
	api.Delete("/:id", middleware.AuthRequired, func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := database.DB.Delete(&models.Book{}, id).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete book"})
		}
		return c.JSON(fiber.Map{"message": "Book deleted"})
	})
}
