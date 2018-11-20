// Package alarmquote represents domain business entities and rules.
// See: "The Clean Architecture" by Uncle Bob
package alarmquote

// ArticleID represents an article ID
type ArticleID string

// Article represents a single article
type Article struct {
	ID          ArticleID
	Name        string
	Description string
	Category    string
	Price       float64
}

// ArticlesService represents a service for managing articles
type ArticlesService interface {
	GetByID(id ArticleID) (*Article, error)
	Add(a *Article) error
}
