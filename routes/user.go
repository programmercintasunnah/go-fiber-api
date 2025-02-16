package routes

import (
	"go-fiber-api/database"
	"go-fiber-api/middleware"
	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/users")
	api.Use(middleware.AuthRequired)

	api.Get("/", func(c *fiber.Ctx) error {
		var users []models.User
		database.DB.Find(&users)
		return c.JSON(users)
	})
}
