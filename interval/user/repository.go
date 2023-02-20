package user

import (
	"context"
	"time"

	"github.com/tunardev/go-api-boilerplate/interval/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	// Create creates a new user.
	Create(user models.User) (models.User, error)

	// GetByEmail gets a user by email.
	GetByEmail(email string) (models.User, error)

	// GetByUsername gets a user by username.
	GetByUsername(username string) (models.User, error)

	// GetByID gets a user by ID.
	GetByID(id string) (models.User, error)
}

type repository struct {
	// collection is the MongoDB collection for the user service.
	collection *mongo.Collection
}

// NewRepository creates a new user repository.
func NewRepository(collection *mongo.Collection) Repository {
	return repository{collection}
}

func (r repository) Create(user models.User) (models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert the user into the database.
	res, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}

	// Set the ID of the user to the ID of the inserted document.
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r repository) GetByEmail(email string) (models.User, error) {
	// Create a new user.
	user := models.User{}

	// Find the user in the database.
	err := r.collection.FindOne(context.TODO(), models.User{Email: email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r repository) GetByUsername(username string) (models.User, error) {
	// Create a new user.
	user := models.User{}

	// Find the user in the database.
	err := r.collection.FindOne(context.TODO(), models.User{Username: username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r repository) GetByID(id string) (models.User, error) {
	// Convert the ID to a primitive.ObjectID.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	// Create a new user.
	user := models.User{}

	// Find the user in the database.
	err = r.collection.FindOne(context.TODO(), models.User{ID: objID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}