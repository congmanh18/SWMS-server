package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App, repository *handlers.Repository) {
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	app.Post("/transaction/create", handler.Repository.CreateTransaction)
	app.Get("/transaction/:id", handler.Repository.ReadTransaction)
	app.Get("/transaction/", handler.Repository.ReadListTransaction)
	app.Put("/transaction/:id", handler.Repository.UpdateTransaction)
	app.Delete("/transaction/:id", handler.Repository.DeleteTransaction)
}
