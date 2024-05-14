package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func AreaRoutes(app *fiber.App, repository *handlers.Repository) {
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	app.Post("/area/create", handler.Repository.CreateArea)
	app.Get("/area/:id", handler.Repository.ReadArea)
	app.Get("/area", handler.Repository.ReadListArea)
	app.Put("/area/:id", handler.Repository.UpdateArea)
	app.Delete("/area/:id", handler.Repository.DeleteArea)
}
