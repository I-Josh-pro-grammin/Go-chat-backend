package Controllers

import (
    "chat/handlers"
	"chat/config"
	"chat/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"time"
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



func Login(c fiber.Ctx) {
   type LoginRequest struct{
	Username string `json:"username"`
   }

   var req LoginRequest;

   if err := c.Bind().Body(&req); err != nil {
	  c.Status(400).JSON(fiber.Map{
		"error": "Invalid input",
	  })
   }
   
   var user models.User;

   if err := config.DB.Where("username = ?", req.Username).Find(&user); err != nil {
	  c.Status(400).JSON(fiber.Map{
		"message": "User doesn't exist!",
	})
   }

   c.Status(200).JSON(user);
   
}