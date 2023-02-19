package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Story is the model for a story.
type Story struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
