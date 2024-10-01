package main

import (
	"context"
	"log"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
	"github.com/newsesame/jblog/database"
	"github.com/newsesame/jblog/router"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {

		// log on terminal
		log.Fatal("Error in loading .env file.")
	}

	// Initialize the DBConn variable
	database.ConnectDB()
}
func main() {

	// Take the DBConnection from the database package
	mongoTestClient := database.DBConn

	defer mongoTestClient.Disconnect(context.Background())

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// A Middleware logger
	app.Use(logger.New())

	// routing
	router.SetupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"DLLM": true})
	// })
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello Wrold")
	// })

	// app.Listen(":8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
