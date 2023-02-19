package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tunardev/go-api-boilerplate/interval/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load config
	cfg, err := config.Load("./config/config.json")
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Mongo_URI))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO()) // Close connection when main() exits
	
	// Build web app
	app := fiber.New()

	// middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go REST API starter kit")
	})

	// Start server
	app.Listen(fmt.Sprintf(":%v", cfg.Port))
}