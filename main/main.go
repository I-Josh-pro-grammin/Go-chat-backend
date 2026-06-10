package main;

import (
    "encoding/json"
    "log"
    "os"
	"fmt"
    "github.com/gofiber/fiber/v3"
    "github.com/pusher/pusher-http-go/v5"
)

// type Message struct{
//     UserID string  `json:"userId"`;
//     RoomID string  `json:"roomId"`;
//     Content string `json:"content"`
// }

// func main() {
//     // Initialize a new Fiber app
//     app := fiber.New()

// 	pusherClient := pusher.Client{
// 		AppID: os.Getenv(PUSHER_APP_ID),
// 		Key: os.Getenv(PUSHER_KEY),
// 		Secret: os.Getenv(PUSHER_SECRET),
// 		Cluster: os.Getenv(PUSHER_CLUSTER),
// 		Secure: true,
// 	}

//     // Define a route for the GET method on the root path '/'
//     app.Post("/api/messages", func(c fiber.Ctx) error {
//         // Send a string response to the client
// 		var msg Message;

// 		// Parse request body as JSON: {"key":"value"}
// 		var raw = c.Body()
//         if err := c.Bind().Body(&msg); err != nil {
//             return c.Status(400).JSON(fiber.Map{
//                 "error": "Invalid request body."
//             })
//         }

// 		if err := json.Unmarshal(raw, &data); err != nil {
// 			return err
// 		}


//         // Trigger the event and capture the response/error
//         if err := pusherClient.Trigger("chat", "message", data); err != nil {
//             log.Printf("Pusher trigger error: %v", err)
//             return c.Status(500).SendString(err.Error())
//         }

//         fmt.Printf("Received message: %s\n", data["message"])
//         return c.JSON([]string{})
//     })
    

//     // Start the server on port 8000
//     log.Fatal(app.Listen(":8000"))
// }

import (

)

function main() {
    godotenv.Load()
    config.connectDB()
    config.initPusher()

    config.DB.AutoMigrate(
        &models.User{},
        &models.Room{},
        &models.Message{}
    )

    app := fiber.New();

    app.use(cors.New());
    routes.Setup(app);

    log.Fatal(app.Listen(os.Getenv("PORT") || "5000"))
}