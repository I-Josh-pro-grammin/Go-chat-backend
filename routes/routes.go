package routes

import (
	"chat/Controllers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	api := app.Group("/api");

	api.Post("/users/register", Controllers.RegisterUser)
	api.Get("/users/:userId", Controllers.GetUserById)
	api.Post("/users/login", Controllers.Login)
    
	api.Post("/messages", Controllers.CreateMessage)

	api.Get("/rooms/:roomId/messages", Controllers.GetRoomMessages)

	// api.Get("/rooms/:messageId", Controllers.GetRoomByMessageId)
}