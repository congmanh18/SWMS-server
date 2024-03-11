package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, uh *handlers.UserHandler) {
	app.Post("/users/register", uh.Register)
	app.Post("/users/login", uh.Login)
	app.Get("/users/:id", uh.ViewInfo)
}
