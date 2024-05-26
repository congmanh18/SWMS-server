package routes

import (
	"smart-waste-system/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func TrashBinRoutes(app *fiber.App, repository *handlers.Repository) {
	handler := &handlers.Handler{Repository: repository}

	app.Post("/trashBin/create", handler.Repository.CreateTrashBin)
	app.Get("/trashBin", handler.Repository.ReadListTrashbin)
	app.Get("/trashBin/:id", handler.Repository.ReadTrashBin)
	app.Delete("/trashBin/:id", handler.Repository.DeleteTrashBin)
	app.Put("/trashBin/:id", handler.Repository.UpdateTrashBin)
	app.Get("ws/trashBin/:id", websocket.New(func(c *websocket.Conn) {
		handler.Repository.WebSocketReadTrashBin(c)
	}))
	app.Put("ws/trashBin/:id", websocket.New(func(c *websocket.Conn) {
		handler.Repository.WebSocketUpdateTrashBin(c)
	}))
}
