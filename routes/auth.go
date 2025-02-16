package routes

import (
	"fmt"
	"go-fiber-api/database"
	"go-fiber-api/middleware"
	"go-fiber-api/models"
	"time"

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

	const maxFailedAttempts = 5
	const lockDuration = 15 * time.Minute

	api.Post("/login", middleware.RateLimiter(), func(c *fiber.Ctx) error {
		var input models.User
		var user models.User

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		database.DB.Where("username = ?", input.Username).First(&user)
		if user.ID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "User not found"})
		}

		fmt.Println("LockedUntil:", user.LockedUntil)
		fmt.Println("Current Time:", time.Now().Unix())

		if user.LockedUntil != 0 && user.LockedUntil > time.Now().Unix() {
			return c.Status(403).JSON(fiber.Map{"error": "Akun dikunci, coba lagi nanti"})
		}

		// Cek password
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
			user.FailedLogins++
			if user.FailedLogins >= maxFailedAttempts {
				user.LockedUntil = time.Now().Add(lockDuration).Unix() // Kunci akun
				fmt.Println("Locking Account Until:", user.LockedUntil)
			}
			err := database.DB.Save(&user).Error
			if err != nil {
				fmt.Println("Error saving user:", err)
			}
			return c.Status(400).JSON(fiber.Map{"error": "Wrong password"})
		}

		// Reset gagal login jika sukses
		user.FailedLogins = 0
		user.LockedUntil = 0

		err := database.DB.Save(&user).Error
		if err != nil {
			fmt.Println("Error saving user:", err)
		}

		token, _ := middleware.GenerateToken(user.Username)
		return c.JSON(fiber.Map{"token": token})
	})
}
