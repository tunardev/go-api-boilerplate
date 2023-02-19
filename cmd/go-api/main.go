package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tunardev/go-api-boilerplate/interval/config"
	"github.com/tunardev/go-api-boilerplate/interval/story"
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
	db := client.Database("project")
	defer func () { // Close connection when main() exits
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	
	// Build web app
	app := fiber.New()

	// initalize middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))
	
	// initalize routes
	story.Handler(app, db.Collection("stories"))

	// Start server
	app.Listen(fmt.Sprintf(":%v", cfg.Port))
}