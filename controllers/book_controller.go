package controllers

import (
	"go-fiber-api/models"
	"go-fiber-api/repositories"
	"go-fiber-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	books, err := repositories.BookRepository.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve books"})
	}
	return c.JSON(books)
}

func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := repositories.BookRepository.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := repositories.BookRepository.Create(&book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := repositories.BookRepository.Update(id, &book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update book"})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repositories.BookRepository.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.JSON(fiber.Map{"message": "Book deleted"})
}

func SearchBooks(c *fiber.Ctx) error {
	var params utils.SearchParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	books, total, err := repositories.BookRepository.Search(params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to search books"})
	}

	response := fiber.Map{
		"data":       books,
		"total":      total,
		"page":       params.Page,
		"pageSize":   params.PageSize,
		"totalPages": (total + int64(params.PageSize) - 1) / int64(params.PageSize),
	}

	return c.JSON(response)
}
