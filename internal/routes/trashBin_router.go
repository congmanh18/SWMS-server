package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func TrashBinRoutes(app *fiber.App, repository *handlers.Repository) {
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	app.Post("/trashBin/create", handler.Repository.CreateTrashBin)
	app.Get("/trashBin/:id", handler.Repository.ReadTrashBin)
	app.Put("/trashBin/:id", handler.Repository.UpdateTrashBin)
	app.Delete("/trashBin/:id", handler.Repository.DeleteTrashBin)
}
