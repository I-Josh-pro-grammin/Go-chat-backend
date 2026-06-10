package routes

import (
	"chat/Controllers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	api := app.Group("/api");
    
	api.Post("/messages", Controllers.CreateMessage)

	api.Get("/rooms/:roomId/messages", Controllers.GetRoomMessages)
}