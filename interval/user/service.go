package user

import (
	"github.com/tunardev/go-api-boilerplate/interval/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// Register creates a new user.
	Register(user models.User) (models.User, error)

	// Login login in a user.
	Login(user models.User) (models.User, error)

	// Me gets the current user.
	Me(id string) (models.User, error)
}

type service struct {
	repo Repository
}

// NewService creates a new user service.
func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Register(user models.User) (models.User, error) {
	// Hash the password.
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(bytes)

	// Create the user in the database.
	user, err = s.repo.Create(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s service) Login(user models.User) (models.User, error) {
	// Get the user from the database.
	userData, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return models.User{}, err
	}

	// Compare the password.
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s service) Me(id string) (models.User, error) {
	// Get the user from the database.
	user, err := s.repo.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
