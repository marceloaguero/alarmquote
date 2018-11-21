package alarmquote

// General errors.
const (
	ErrUnauthorized = Error("unauthorized")
)

// Articles errors
const (
	// ErrArticleNotFound means that the article could not be found in the repository
	ErrArticleNotFound = Error("article not found")

	// ErrArticleExists means that the article exists in the repository
	ErrArticleExists = Error("article already exists")

	// ErrArticleIDRequired means that article require an ID
	ErrArticleIDRequired = Error("article id required")

	// ErrArticleNameRequired means that article require a name
	ErrArticleNameRequired = Error("article name required")

	// ErrArticleCategoryRequired means that the article require a category
	ErrArticleCategoryRequired = Error("article category required")
)

// Error represents an domain error
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }