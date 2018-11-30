package category

import (
	"github.com/marceloaguero/alarmquote"
	"github.com/pkg/errors"
)

// Ensure CategoryService implements alarmquote.CategoryService.
var _ alarmquote.CategoryService = &Service{}

// Service provides the category service.
type Service struct {
	repo Repository
}

// Repository models the concrete data repository (memory, cache, db, etc).
type Repository interface {
	Retrieve(n alarmquote.CategoryName) (*alarmquote.Category, error)
	Insert(c alarmquote.Category) error
	//	Modify(n alarmquote.CategoryName, c alarmquote.Category) error
	//	Delete(n alarmquote.CategoryName) error
}

// NewService returns a usable service, wrapping a repository.
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// Category retrieve a category from the repository, given it's name.
func (s *Service) Category(n alarmquote.CategoryName) (*alarmquote.Category, error) {
	return s.repo.Retrieve(n)
}

// Add adds a category to the service repository.
func (s *Service) Add(c alarmquote.Category) error {
	if err := validate(c); err != nil {
		return err
	}

	_, err := s.Category(c.Name)
	switch err {
	case alarmquote.ErrCategoryNotFound:
		if err := s.repo.Insert(c); err != nil {
			return errors.Wrap(err, "error adding a new category")
		}

	case nil:
		return alarmquote.ErrCategoryExists

	default:
		return err
	}

	return nil
}

// validate performs basic category validation.
func validate(c alarmquote.Category) error {
	if c.Name == "" {
		return alarmquote.ErrCategoryNameRequired
	}

	return nil
}
