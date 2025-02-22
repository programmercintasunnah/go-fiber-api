package controllers

import (
	"go-fiber-api/repositories"
	"go-fiber-api/utils"

	"github.com/gofiber/fiber/v2"
)

// Generic handler for GetAll
func GetAll[T any](repo *repositories.BaseRepository[T], c *fiber.Ctx) error {
	items, err := repo.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve items"})
	}
	return c.JSON(items)
}

// Generic handler for GetByID
func GetByID[T any](repo *repositories.BaseRepository[T], c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := repo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}
	return c.JSON(item)
}

// Generic handler for Create
func Create[T any](repo *repositories.BaseRepository[T], c *fiber.Ctx) error {
	var item T
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := repo.Create(&item); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create item"})
	}

	return c.Status(201).JSON(item)
}

// Generic handler for Update
func Update[T any](repo *repositories.BaseRepository[T], c *fiber.Ctx) error {
	id := c.Params("id")
	var item T
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := repo.Update(id, &item); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update item"})
	}

	return c.JSON(item)
}

// Generic handler for Delete
func Delete[T any](repo *repositories.BaseRepository[T], c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete item"})
	}

	return c.JSON(fiber.Map{"message": "Item deleted"})
}

// Generic handler for Search
func Search[T any](repo *repositories.BaseRepository[T], c *fiber.Ctx) error {
	var params utils.SearchParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	items, total, err := repo.Search(params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to search items"})
	}

	response := fiber.Map{
		"data":       items,
		"total":      total,
		"page":       params.Page,
		"pageSize":   params.PageSize,
		"totalPages": (total + int64(params.PageSize) - 1) / int64(params.PageSize),
	}

	return c.JSON(response)
}
