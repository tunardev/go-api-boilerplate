package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tunardev/go-api-boilerplate/interval/models"
	"github.com/tunardev/go-api-boilerplate/pkg/errors"
)

type routes struct {
	service Service
	store *session.Store
}

func (r routes) register(c *fiber.Ctx) error {
	// Create a new user.
	user := models.User{}

	// Parse the request body into the user.
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	// Validate the user.
	if err := user.ValidateRegister(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	// Register the user.
	user, status, err := r.service.Register(user)
	if err != nil {
		return c.Status(status).JSON(errors.InternalServerError(err.Error()))
	}

	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}

	// Set the user ID in the session.
	sess.Set("userID", user.ID.Hex())
	if err = sess.Save(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}
	
	// Return the user.
	return c.Status(status).JSON(fiber.Map{
		"data": user,
		"status": status,
	})
}

func (r routes) login(c *fiber.Ctx) error {
	// Create a new user.
	user := models.User{}

	// Parse the request body into the user.
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	// Validate the user.
	if err := user.ValidateLogin(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	// Login the user.
	user, status, err := r.service.Login(user)
	if err != nil {
		return c.Status(status).JSON(errors.InternalServerError(err.Error()))
	}

	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}

	// Set the user ID in the session.
	sess.Set("userID", user.ID.Hex())
	if err = sess.Save(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the user.
	return c.Status(status).JSON(fiber.Map{
		"data": user,
		"status": status,
	})
}

func (r routes) me(c *fiber.Ctx) error {
	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}

	// Get the user ID from the session.
	userID := sess.Get("userID")
	if userID == nil {
		return c.Status(http.StatusUnauthorized).JSON(errors.Unauthorized(""))
	}

	// Get the user from the database.
	user, status, err := r.service.Me(userID.(string))
	if err != nil {
		return c.Status(status).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the user.
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": user,
		"status": http.StatusOK,
	})
}

func (r routes) logout(c *fiber.Ctx) error {
	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}

	// Delete the user ID from the session.
	sess.Delete("userID")
	if err = sess.Save(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the user.
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out.",
		"status": http.StatusOK,
	})
}
