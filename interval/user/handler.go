package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is the handler for the user service.
func Handler(app *fiber.App, collection *mongo.Collection, store *session.Store) {
	res := routes{
	 	service: NewService(NewRepository(collection)),
		store: store,
	}
	r := app.Group("/user")

	// user routes
	r.Post("/register", res.register)
	r.Post("/login", res.login)
	r.Get("/me", res.me)
	r.Get("/logout", res.logout)
}
