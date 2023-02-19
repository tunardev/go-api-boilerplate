package story

import (
	"context"
	"time"

	"github.com/tunardev/go-api-boilerplate/interval/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	// Create creates a new story.
	Create(story models.Story) (models.Story, error)

	// Get gets a story by ID.
	Get(id string) (models.Story, error)

	// Delete deletes a story by ID.
	Update(id string, story models.Story) (models.Story, error)
}

type repository struct {
	// collection is the MongoDB collection for the story service.
	collection *mongo.Collection
}

// NewRepository creates a new story repository.
func NewRepository(collection *mongo.Collection) Repository {
	return repository{collection}
}

func (r repository) Create(story models.Story) (models.Story, error) {
	story.CreatedAt = time.Now()
	story.UpdatedAt = time.Now()

	// Insert the story into the database.
	res, err := r.collection.InsertOne(context.TODO(), story)
	if err != nil {
		return models.Story{}, err
	}

	// Set the ID of the story to the ID of the inserted document.
	story.ID = res.InsertedID.(primitive.ObjectID)
	return story, nil
}

func (r repository) Get(id string) (models.Story, error) {
	// Convert the ID to a primitive.ObjectID.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Story{}, err
	}

	// Create a new story.
	story := models.Story{}

	// Find the story in the database.
	err = r.collection.FindOne(context.TODO(), models.Story{ID: objID}).Decode(&story)
	if err != nil {
		return models.Story{}, err
	}

	return story, nil
}

func (r repository) Update(id string, story models.Story) (models.Story, error) {
	// Convert the ID to a primitive.ObjectID.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Story{}, err
	}
	story.ID = objID
	story.UpdatedAt = time.Now()

	// Update the story in the database.
	_, err = r.collection.UpdateOne(context.TODO(), models.Story{ID: objID}, story)
	if err != nil {
		return models.Story{}, err
	}

	return story, nil
}
