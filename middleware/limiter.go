package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// Rate limiter untuk login
func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        5,               // Maksimal 5 request
		Expiration: 5 * time.Minute, // Dalam 5 menit
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Gunakan IP sebagai kunci
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{"error": "Terlalu banyak percobaan, coba lagi nanti"})
		},
	})
}
