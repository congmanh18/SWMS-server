package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, repository *handlers.Repository) {
	handler := &handlers.Handler{Repository: repository} // Assuming Handler is defined in the handlers package

	app.Post("/users/register", handler.Repository.Register)
	app.Post("/users/login", handler.Repository.Login)
	app.Get("/users/:id", handler.Repository.ViewInfo)
}
