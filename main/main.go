package main;

import (
    "Backend/config"
    "Backend/models"
    "log"
    "os"
	"fmt"
    "github.com/joho/godotenv"
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
    godotenv.Load()
    config.connectDB()
    config.initPusher()

    config.DB.AutoMigrate(
        &models.User{},
        &models.Room{},
        &models.Message{},
    )

    app := fiber.New();

    app.use(cors.New());
    routes.Setup(app);

    log.Fatal(app.Listen(os.Getenv("PORT") || "5000"))
}