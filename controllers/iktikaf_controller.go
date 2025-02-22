package controllers

import (
	"go-fiber-api/repositories"

	"github.com/gofiber/fiber/v2"
)

// Menggunakan generic handlers
func GetAllIktikafs(c *fiber.Ctx) error {
	return GetAll(repositories.IktikafRepository, c)
}

func GetIktikafByID(c *fiber.Ctx) error {
	return GetByID(repositories.IktikafRepository, c)
}

func CreateIktikaf(c *fiber.Ctx) error {
	return Create(repositories.IktikafRepository, c)
}

func UpdateIktikaf(c *fiber.Ctx) error {
	return Update(repositories.IktikafRepository, c)
}

func DeleteIktikaf(c *fiber.Ctx) error {
	return Delete(repositories.IktikafRepository, c)
}

func SearchIktikafs(c *fiber.Ctx) error {
	return Search(repositories.IktikafRepository, c)
}
