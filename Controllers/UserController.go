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

   if req.Username == "" {
	   return c.Status(400).JSON(fiber.Map{
		"error": "Invalid input",
	   })
   }

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

func findRoomUsers(c fiber.Ctx) error {
   roomId := c.Params("roomId")

   var users []models.User;

   config.DB.
	Joins("Message").
	Where("Message.room_id = ?", roomId).
	Find(&users);

	return c.JSON(users)
}
