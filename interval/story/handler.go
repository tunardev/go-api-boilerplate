package story

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is the handler for the story service.
func Handler(app *fiber.App, collection *mongo.Collection) {
	// res := routes{
	// 	service: NewService(NewRepository(collection)),
	// }
	// r := app.Group("/story")

	// story routes
	// r.Get("/", ...)
}