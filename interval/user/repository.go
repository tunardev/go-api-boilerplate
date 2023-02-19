package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
}

type repository struct {
	// collection is the MongoDB collection for the user service.
	collection *mongo.Collection
}

// NewRepository creates a new user repository.
func NewRepository(collection *mongo.Collection) Repository {
	return repository{collection}
}
