package models

import (
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Validate validates the user.
func (u User) Validate() (string, bool) {
	usernameRegex := regexp.MustCompile(`/^\w+$/i`)
	if !usernameRegex.MatchString(u.Username) {
		return "Username must be a nonempty alphanumeric string.", false
	}

	passwordRegex := regexp.MustCompile(`/^\S+$/`)
	if !passwordRegex.MatchString(u.Password) {
		return "Password must be a nonempty string.", false
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return "Email must be a valid email address.", false
	}

	return "", true
}

