package routes

import (
	"go-fiber-api/database"
	"go-fiber-api/middleware"
	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/auth")

	api.Post("/register", func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Hash password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		database.DB.Create(&user)
		return c.JSON(fiber.Map{"message": "User created successfully"})
	})

	api.Post("/login", func(c *fiber.Ctx) error {
		var input models.User
		var user models.User

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		database.DB.Where("username = ?", input.Username).First(&user)
		if user.ID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "User not found"})
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Wrong password"})
		}

		token, _ := middleware.GenerateToken(user.Username)
		return c.JSON(fiber.Map{"token": token})
	})
}
