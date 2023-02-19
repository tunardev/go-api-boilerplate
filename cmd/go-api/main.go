package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tunardev/go-api-boilerplate/interval/config"
)

func main() {
	// Load config
	cfg, err := config.Load("./config/config.json")
	if err != nil {
		panic(err)
	}
	
	// Build web app
	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go REST API starter kit")
	})

	// Start server
	app.Listen(fmt.Sprintf(":%v", cfg.Port))
}