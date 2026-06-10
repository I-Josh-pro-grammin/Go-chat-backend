package main;

import (
    "chat/config"
    "chat/models"
    "log"
    "os"
    "chat/routes"
    "github.com/joho/godotenv"
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
    godotenv.Load()
    config.ConnectDB()
    config.InitPusher()

    config.DB.AutoMigrate(
        &models.User{},
        &models.Room{},
        &models.Message{},
    )

    app := fiber.New();

    app.Use(cors.New());
    routes.Setup(app);

    log.Fatal(app.Listen(os.Getenv("PORT")))
}