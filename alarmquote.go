// Package alarmquote represents domain business entities and rules.
// See: "The Clean Architecture" by Uncle Bob.
package alarmquote

// ArticleID represents an article ID.
type ArticleID string

// Article represents a single article.
type Article struct {
	ID          ArticleID
	Name        string
	Description string
	Category    string
	Price       float64
}

// ArticleService represents a service for managing articles.
type ArticleService interface {
	Article(id ArticleID) (*Article, error)
	Add(a Article) error
	Edit(id ArticleID, a Article) error
	Delete(id ArticleID) error
}

// CategoryName repesents a category's name
type CategoryName string

// Category represents an article's category
type Category struct {
	Name CategoryName
}

// CategoryService is a service for managing articles' categories
type CategoryService interface {
	Category(n CategoryName) (*Category, error)
	Add(c Category) error
	Edit(n CategoryName, c Category) error
	Delete(n CategoryName) error
}
