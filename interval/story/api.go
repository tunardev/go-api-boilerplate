package story

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tunardev/go-api-boilerplate/interval/models"
	"github.com/tunardev/go-api-boilerplate/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type routes struct {
	service Service
}

func (r routes) create(c *fiber.Ctx) error {
	// Create a new story.
	story := models.Story{}

	// Parse the request body into the story.
	err := c.BodyParser(&story)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	// Validate the story.
	if ok := story.Validate(); !ok {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest("invalid story"))
	}

	// Get the user ID from the context.
	userID := c.Locals("userID")
	if userID == nil {
		return c.Status(http.StatusUnauthorized).JSON(errors.Unauthorized(""))
	}

	// Convert the user ID to an ObjectID.
	objID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.InternalServerError(err.Error()))
	}
	story.UserID = objID

	// Create the story in the database.
	story, status, err := r.service.Create(story)
	if err != nil {
		return c.Status(status).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the story.
	return c.Status(status).JSON(fiber.Map{
		"data": story,
		"status": status,
	})
}

func (r routes) get(c *fiber.Ctx) error {
	// Get the ID from the request parameters.
	id := c.Params("id")

	// Get the story from the database.
	story, status, err := r.service.Get(id)
	if err != nil {
		return c.Status(status).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the story.
	return c.Status(status).JSON(fiber.Map{
		"data": story,
		"status": status,
	})
}

func (r routes) update(c *fiber.Ctx) error {
	// Get the ID from the request parameters.
	id := c.Params("id")

	// Create a new story.
	story := models.Story{}

	// Parse the request body into the story.
	err := c.BodyParser(&story)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest(err.Error()))
	}

	// Validate the story.
	if ok := story.Validate(); !ok {
		return c.Status(http.StatusBadRequest).JSON(errors.BadRequest("invalid story"))
	}

	// Get the user ID from the context.
	userID := c.Locals("userID")
	if userID == nil {
		return c.Status(http.StatusUnauthorized).JSON(errors.Unauthorized(""))
	}

	// Update the story in the database.
	story, status, err := r.service.Update(id, story)
	if err != nil {
		return c.Status(status).JSON(errors.InternalServerError(err.Error()))
	}

	// Return the story.
	return c.Status(status).JSON(fiber.Map{
		"data": story,
		"status": status,
	})
}