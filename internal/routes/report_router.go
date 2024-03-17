package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

// /test commit
func ReportRoutes(app *fiber.App, repository *handlers.Repository) {
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	app.Post("/report/create", handler.Repository.CreateReport)
	app.Get("/report/:id", handler.Repository.ReadReport)
	app.Put("/report/:id", handler.Repository.UpdateReport)
	app.Delete("/report/:id", handler.Repository.DeleteReport)
}
