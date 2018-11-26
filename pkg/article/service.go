package article

import (
	"github.com/marceloaguero/alarmquote"
	"github.com/pkg/errors"
)

// Ensure ArticleService implements alarmquote.ArticleService.
var _ alarmquote.ArticleService = &Service{}

// Service provides the article service.
type Service struct {
	repo Repository
}

// Repository models the concrete data repository (memory, cache, db, etc).
type Repository interface {
	Retrieve(id alarmquote.ArticleID) (*alarmquote.Article, error)
	Insert(a alarmquote.Article) error
	Modify(id alarmquote.ArticleID, a alarmquote.Article) error
	Delete(id alarmquote.ArticleID) error
}

// NewService returns a usable service, wrapping a repository.
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// Article retrieve an article from the repository, given it's ID.
func (s *Service) Article(id alarmquote.ArticleID) (*alarmquote.Article, error) {
	return s.repo.Retrieve(id)
}

// Add adds an article to the service repository.
func (s *Service) Add(a alarmquote.Article) error {
	if err := validate(a); err != nil {
		return err
	}

	_, err := s.Article(a.ID)
	switch err {
	case alarmquote.ErrArticleNotFound:
		if err := s.repo.Insert(a); err != nil {
			return errors.Wrap(err, "error adding a new article")
		}

	case nil:
		return alarmquote.ErrArticleExists

	default:
		return err
	}

	return nil
}

// Edit permits article's modifications.
func (s *Service) Edit(id alarmquote.ArticleID, a alarmquote.Article) error {
	if err := validate(a); err != nil {
		return err
	}

	if id != a.ID {
		return alarmquote.ErrChangeIDForbidden
	}

	if _, err := s.Article(id); err != nil {
		return errors.Wrap(err, "error retrieving when editing an article")
	}

	if err := s.repo.Modify(id, a); err != nil {
		return errors.Wrap(err, "error editing an article")
	}

	return nil
}

func (s *Service) Delete(id alarmquote.ArticleID) error {
	if _, err := s.Article(id); err != nil {
		return alarmquote.ErrArticleNotFound
	}

	if err := s.repo.Delete(id); err != nil {
		return errors.Wrap(err, "error deleting an article")
	}

	return nil
}

// validate performs basic article validation.
func validate(a alarmquote.Article) error {
	if a.ID == "" {
		return alarmquote.ErrArticleIDRequired
	}

	if a.Name == "" {
		return alarmquote.ErrArticleNameRequired
	}

	if a.Category == "" {
		return alarmquote.ErrArticleCategoryRequired
	}

	return nil
}
