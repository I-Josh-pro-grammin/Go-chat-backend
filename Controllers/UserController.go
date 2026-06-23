package Controllers

import (
    "chat/handlers"
	"chat/config"
	"chat/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"time"
	"errors"
	"gorm.io/gorm"
)

func RegisterUser(c fiber.Ctx) error {
   var req handlers.UserRequest

   if err := c.Bind().Body(&req);  err != nil {
	return c.Status(400).JSON(fiber.Map{
		"error": "invalid request",
	})
   }

   user := models.User{
	ID: uuid.New(),
	Username: req.Username,
	Created_At: time.Now(),
   }

   if err := config.DB.Create(&user).Error; err != nil {
	return c.Status(500).JSON(fiber.Map{
		"error": "User creation failed",
	})
   }

   room := models.Room{
	ID: uuid.New(),
	Name: "#" + req.Username,
	Username: req.Username,
	Created_At: time.Now(),
   }

   if err := config.DB.Create(&room).Error; err != nil {
	return c.Status(500).JSON(fiber.Map{
		"error": "Room creation failed",
	})
   }

   return c.JSON(user)
}



func GetUserById(c fiber.Ctx) error {
    userId := c.Params("userId");
	var user models.User;

	config.DB.Where("id = ?", userId).
	Find(&user);

	return c.JSON(user);
}



func Login(c fiber.Ctx) error { // 1. Change signature to return error
   type LoginRequest struct{
	  Username string `json:"username"`
   }

   var req LoginRequest;

   // 2. Use '!= nil' and 'return' to halt execution on bad input
   if err := c.Bind().Body(&req); err != nil {
	  return c.Status(400).JSON(fiber.Map{
		"error": "Invalid input",
	  })
   }
   
   var user models.User;

   // 3. Use 'First' and check for ErrRecordNotFound
   // (Make sure to import "errors" and "gorm.io/gorm" in your UserController imports)
   if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
      if errors.Is(err, gorm.ErrRecordNotFound) {
          return c.Status(404).JSON(fiber.Map{
              "message": "User doesn't exist!",
          })
      }
      return c.Status(500).JSON(fiber.Map{
          "error": "Database error",
      })
   }

   // 4. Return the user successfully
   return c.Status(200).JSON(user)
}

func GetRoomUsers(c fiber.Ctx) error {
   roomId := c.Params("roomId")

   if _, err := uuid.Parse(roomId); err != nil {
       return c.Status(400).JSON(fiber.Map{
           "error": "Invalid room ID format",
       })
   }

   var users []models.User;

   if err := config.DB.
	Where("id::text IN (SELECT user_id::text FROM messages WHERE room_id::text = ?)", roomId).
	Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(users);
}
