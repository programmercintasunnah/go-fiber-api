package controllers

import (
	"go-fiber-api/repositories"

	"github.com/gofiber/fiber/v2"
)

// Menggunakan generic handlers
func GetAllBooks(c *fiber.Ctx) error {
	return GetAll(repositories.BookRepository, c)
}

func GetBookByID(c *fiber.Ctx) error {
	return GetByID(repositories.BookRepository, c)
}

func CreateBook(c *fiber.Ctx) error {
	return Create(repositories.BookRepository, c)
}

func UpdateBook(c *fiber.Ctx) error {
	return Update(repositories.BookRepository, c)
}

func DeleteBook(c *fiber.Ctx) error {
	return Delete(repositories.BookRepository, c)
}

func SearchBooks(c *fiber.Ctx) error {
	return Search(repositories.BookRepository, c)
}
