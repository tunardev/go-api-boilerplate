package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tunardev/go-api-boilerplate/pkg/errors"
)

func IsAuth(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the session.
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
		}

		// Get the user ID from the session.
		userID := sess.Get("userID")
		if userID != nil {
			// Set the user ID in the context.
			c.Locals("userID", userID)
			return c.Next()
		}

		// Return unauthorized.
		return c.Status(http.StatusUnauthorized).JSON(errors.Unauthorized(""))
	}
}