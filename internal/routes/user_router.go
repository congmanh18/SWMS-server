package routes

import (
	"smart-waste-system/internal/app/repository"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, uh *repository.Repository) {

	handler := &repository.Handler{Repository: uh}

	app.Post("/users/register", handler.Repository.)
	app.Post("/users/login", handler.Repository.)
	app.Get("/users/:id", handler.Repository.)
}
