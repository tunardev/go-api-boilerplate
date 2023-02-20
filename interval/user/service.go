package user

import (
	"errors"
	"net/http"

	"github.com/tunardev/go-api-boilerplate/interval/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// Register creates a new user.
	Register(user models.User) (models.User, int, error)

	// Login login in a user.
	Login(user models.User) (models.User, int, error)

	// Me gets the current user.
	Me(id string) (models.User, int, error)
}

type service struct {
	repo Repository
}

// NewService creates a new user service.
func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Register(user models.User) (models.User, int, error) {
	// Check if the username or email already exists.
	_, err := s.repo.GetByUsername(user.Username)
	if err != nil {
		return models.User{}, http.StatusBadRequest, errors.New("username already exists")
	}
	_, err = s.repo.GetByEmail(user.Email)
	if err != nil {
		return models.User{}, http.StatusBadRequest, errors.New("email already exists")
	}

	// Hash the password.
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}
	user.Password = string(bytes)

	// Create the user in the database.
	user, err = s.repo.Create(user)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	return user, http.StatusCreated, nil
}

func (s service) Login(user models.User) (models.User, int, error) {
	// Get the user from the database.
	userData, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return models.User{}, http.StatusBadRequest, err
	}

	// Compare the password.
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (s service) Me(id string) (models.User, int, error) {
	// Get the user from the database.
	user, err := s.repo.GetByID(id)
	if err != nil {
		return models.User{}, http.StatusBadRequest, err
	}

	return user, http.StatusOK, nil
}
