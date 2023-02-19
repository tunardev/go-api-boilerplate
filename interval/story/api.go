package story

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tunardev/go-api-boilerplate/interval/models"
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
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"status": 400,
		})
	}

	if ok := story.Validate(); !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid story",
			"status": 400,
		})
	}

	// Create the story in the database.
	story, err = r.service.Create(story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
			"status": 500,
		})
	}

	// Return the story.
	return c.Status(201).JSON(fiber.Map{
		"story": story,
		"status": 201,
	})
}

func (r routes) get(c *fiber.Ctx) error {
	// Get the ID from the request parameters.
	id := c.Params("id")

	// Get the story from the database.
	story, err := r.service.Get(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
			"status": 500,
		})
	}

	// Return the story.
	return c.Status(200).JSON(fiber.Map{
		"story": story,
		"status": 200,
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
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"status": 400,
		})
	}

	if ok := story.Validate(); !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid story",
			"status": 400,
		})
	}

	// Update the story in the database.
	story, err = r.service.Update(id, story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
			"status": 500,
		})
	}

	// Return the story.
	return c.Status(200).JSON(fiber.Map{
		"data": story,
		"status": 200,
	})
}