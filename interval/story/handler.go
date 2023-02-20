package story

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tunardev/go-api-boilerplate/interval/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is the handler for the story service.
func Handler(app *fiber.App, collection *mongo.Collection, store *session.Store) {
	res := routes{
		service: NewService(NewRepository(collection)),
	}
	r := app.Group("/story")

	// story routes
	r.Get("/:id", res.get)

	// protected routes
	r.Use(middleware.IsAuth(store))

	r.Post("/", res.create)
	r.Patch("/:id", res.update)
}