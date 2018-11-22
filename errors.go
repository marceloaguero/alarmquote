package alarmquote

import (
	"github.com/pkg/errors"
)

// General errors.
var (
	ErrUnauthorized = errors.New("unauthorized")
)

// Articles errors
var (
	// ErrArticleNotFound means that the article could not be found in the repository
	ErrArticleNotFound = errors.New("article not found")

	// ErrArticleExists means that the article exists in the repository
	ErrArticleExists = errors.New("article already exists")

	// ErrArticleIDRequired means that article require an ID
	ErrArticleIDRequired = errors.New("article id required")

	// ErrArticleNameRequired means that article require a name
	ErrArticleNameRequired = errors.New("article name required")

	// ErrArticleCategoryRequired means that the article require a category
	ErrArticleCategoryRequired = errors.New("article category required")

	// ErrChangeIDForbidden means that the article ID modification is not allowed
	ErrChangeIDForbidden = errors.New("article ID edition forbidden")
)
