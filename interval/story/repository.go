package story

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
}

type repository struct {
	// db is the MongoDB collection for the story service.
	collection *mongo.Collection
}

// NewRepository creates a new story repository.
func NewRepository(collection *mongo.Collection) Repository {
	return repository{collection}
}


