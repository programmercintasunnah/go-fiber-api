package routes

import (
	"go-fiber-api/controllers"
	"go-fiber-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func IktikafRoutes(app *fiber.App) {
	api := app.Group("/iktikaf")

	api.Get("/", middleware.AuthRequired, controllers.GetAllIktikafs)
	api.Get("/:id", middleware.AuthRequired, controllers.GetIktikafByID)
	api.Post("/", middleware.AuthRequired, controllers.CreateIktikaf)
	api.Put("/:id", middleware.AuthRequired, controllers.UpdateIktikaf)
	api.Delete("/:id", middleware.AuthRequired, controllers.DeleteIktikaf)
	api.Post("/search", middleware.AuthRequired, controllers.SearchIktikafs)
}
