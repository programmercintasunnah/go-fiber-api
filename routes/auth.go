package routes

import (
	"go-fiber-api/controllers"
	"go-fiber-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/auth")

	api.Post("/register", controllers.Register)
	api.Post("/login", middleware.RateLimiter(), controllers.Login)
}
