package routes

import (
	"go-fiber-api/controllers"
	"go-fiber-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	api := app.Group("/books")

	api.Get("/", middleware.AuthRequired, controllers.GetAllBooks)
	api.Get("/:id", middleware.AuthRequired, controllers.GetBookByID)
	api.Post("/", middleware.AuthRequired, controllers.CreateBook)
	api.Put("/:id", middleware.AuthRequired, controllers.UpdateBook)
	api.Delete("/:id", middleware.AuthRequired, controllers.DeleteBook)
	api.Post("/search", middleware.AuthRequired, controllers.SearchBooks)
}
