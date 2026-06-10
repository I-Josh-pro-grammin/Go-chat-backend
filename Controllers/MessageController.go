package Controllers

import (
    "fmt"

	"github.com/google/uuid"
    "chat/handlers"
	"chat/config"
	"chat/models"

	"github.com/gofiber/fiber/v3"
) 

func CreateMessage(c fiber.Ctx) error {
   var req handlers.CreateMessageRequest;

   if err := c.Bind().Body(&req); err != nil {
	  return c.Status(400).JSON(fiber.Map{
		"error": "Invalid input",
	})
   }

   msg := models.Message{
	  ID: uuid.New(),
	  Content: req.Content,
	  UserID: uuid.MustParse(req.UserID),
	  RoomID: uuid.MustParse(req.RoomID),
   }

   if err := config.DB.Create(&msg).Error; err != nil {
	  return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	  })
   }

   channel := fmt.Sprintf("room-%s", req.RoomID);

   config.PusherClient.Trigger(
	channel,
	"message",
	msg,
   )

   return c.JSON(msg);
}

func GetRoomMessages(c fiber.Ctx) error {
	roomId := c.Params("roomId");

    var messages []models.Message

	config.DB.
	Where("room_id = ?", roomId).
	Order("created_at asc").
	Find(&messages)

	return c.JSON(messages)
}