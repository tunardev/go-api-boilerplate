package user

import (
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
		return c.Status(400).JSON(errors.BadRequest(err.Error()))
	}

	// Validate the user.
	if err := user.ValidateRegister(); err != nil {
		return c.Status(400).JSON(errors.BadRequest(err.Error()))
	}

	// Register the user.
	user, err := r.service.Register(user)
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Set the user ID in the session.
	sess.Set("userID", user.ID.Hex())
	
	// Return the user.
	return c.Status(201).JSON(fiber.Map{
		"data": user,
		"status": 201,
	})
}

func (r routes) login(c *fiber.Ctx) error {
	// Create a new user.
	user := models.User{}

	// Parse the request body into the user.
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(errors.BadRequest(err.Error()))
	}

	// Validate the user.
	if err := user.ValidateLogin(); err != nil {
		return c.Status(400).JSON(errors.BadRequest(err.Error()))
	}

	// Login the user.
	user, err := r.service.Login(user)
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Set the user ID in the session.
	sess.Set("userID", user.ID.Hex())

	// Return the user.
	return c.Status(200).JSON(fiber.Map{
		"data": user,
		"status": 200,
	})
}

func (r routes) me(c *fiber.Ctx) error {
	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Get the user ID from the session.
	userID := sess.Get("userID")
	if userID == nil {
		return c.Status(401).JSON(errors.Unauthorized(""))
	}

	// Get the user from the database.
	user, err := r.service.Me(userID.(string))
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the user.
	return c.Status(200).JSON(fiber.Map{
		"data": user,
		"status": 200,
	})
}

func (r routes) logout(c *fiber.Ctx) error {
	// Get the session.
	sess, err := r.store.Get(c)
	if err != nil {
		return c.Status(500).JSON(errors.InternalServerError(err.Error()))
	}

	// Delete the user ID from the session.
	sess.Delete("userID")

	// Return the user.
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully logged out.",
		"status": 200,
	})
}
