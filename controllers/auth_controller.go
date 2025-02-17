package controllers

import (
	"fmt"
	"go-fiber-api/middleware"
	"go-fiber-api/models"
	"go-fiber-api/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

const maxFailedAttempts = 5
const lockDuration = 15 * time.Minute

func Register(c *fiber.Ctx) error {
	var input RegisterUserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validasi password minimal 8 karakter
	if len(input.Password) < 8 {
		return c.Status(400).JSON(fiber.Map{"error": "Password harus minimal 8 karakter"})
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	// Mapping ke model User untuk disimpan ke DB
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := repositories.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}

func Login(c *fiber.Ctx) error {
	var input LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Cek apakah user ada di database
	user, err := repositories.FindUserByUsername(input.Username)
	if err != nil || user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	// Cek apakah akun terkunci
	if user.LockedUntil != 0 && user.LockedUntil > time.Now().Unix() {
		return c.Status(403).JSON(fiber.Map{"error": "Akun dikunci, coba lagi nanti"})
	}

	// Cek password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		user.FailedLogins++
		if user.FailedLogins >= maxFailedAttempts {
			user.LockedUntil = time.Now().Add(lockDuration).Unix()
			fmt.Println("Locking Account Until:", user.LockedUntil)
		}
		repositories.UpdateUser(user)
		return c.Status(400).JSON(fiber.Map{"error": "Wrong password"})
	}

	// Reset gagal login jika sukses
	user.FailedLogins = 0
	user.LockedUntil = 0
	repositories.UpdateUser(user)

	token, _ := middleware.GenerateToken(user.Username)
	return c.JSON(fiber.Map{"token": token})
}
