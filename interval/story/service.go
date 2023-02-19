package story

type Service interface {
}

type service struct {
	repo Repository
}

// NewService creates a new story service.
func NewService(repo Repository) Service {
	return service{repo}
}
