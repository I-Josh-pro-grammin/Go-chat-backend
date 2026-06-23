package Controllers

import (
	// "chat/config"
	"chat/config"
	"chat/models"

	"github.com/gofiber/fiber/v3"
)

func GetRoomByUsername(c fiber.Ctx) error {
	username := c.Params("username");

	var room models.Room;
	
	if(username == "") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := config.DB.Where("username = ?", username).First(&room).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Room not found",
		})
	}

	return c.JSON(room);
}