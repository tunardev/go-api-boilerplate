package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Story is the record for a story.
type Story struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Title string `bson:"title,omitempty" json:"title,omitempty"`
	Text string `bson:"text,omitempty" json:"text,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Validate validates the story.
func (s Story) Validate() bool {
	if (s.Title == "") || (s.Text == "") {
		return false
	}
	
	return true
}