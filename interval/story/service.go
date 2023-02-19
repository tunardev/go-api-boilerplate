package story

import "github.com/tunardev/go-api-boilerplate/interval/models"

type Service interface {
	// Create creates a new story.
	Create(story models.Story) (models.Story, error)

	// Get gets a story by ID.
	Get(id string) (models.Story, error)

	// Update updates a story by ID.
	Update(id string, story models.Story) (models.Story, error)
}

type service struct {
	repo Repository
}

// NewService creates a new story service.
func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Create(story models.Story) (models.Story, error) {
	// Create the story in the database.
	story, err := s.repo.Create(story)
	if err != nil {
		return models.Story{}, err
	}

	return story, nil
}

func (s service) Get(id string) (models.Story, error) {
	// Get the story from the database.
	story, err := s.repo.Get(id)
	if err != nil {
		return models.Story{}, err
	}

	return story, nil
}

func (s service) Update(id string, story models.Story) (models.Story, error) {
	// Update the story in the database.
	story, err := s.repo.Update(id, story)
	if err != nil {
		return models.Story{}, err
	}

	return story, nil
}
