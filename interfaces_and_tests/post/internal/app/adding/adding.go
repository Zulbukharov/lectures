package adding

import (
	"errors"
	"post/internal/models"
	"post/internal/repository"
)

var (
	ErrFailedToAdd  = errors.New("failed to add post")
	ErrInvalidInput = errors.New("invalid input")
)

// Service provides Post adding operations.
type Service interface {
	AddPost(models.PostInsert) error
}

// Repository provides access to Post repository.
type Repository interface {
	Put(*models.PostInsert) error
}

type service struct {
	tR Repository
}

// New creates adding service with the necessary dependencies
func New(r Repository) Service {
	return &service{r}
}

// AddPost adds the given Post to the database
func (s *service) AddPost(u models.PostInsert) error {
	if u.AuthorID == 0 || u.Content == "" {
		return ErrInvalidInput
	}
	err := s.tR.Put(&u)
	if errors.Is(err, repository.ErrFailedToAdd) {
		return ErrFailedToAdd
	}
	return err
}
