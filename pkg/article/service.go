package article

import (
	"github.com/marceloaguero/alarmquote"
)

// Ensure ArticleService implements alarmquote.ArticleService
var _ alarmquote.ArticleService = &ArticleService{}

type ArticleService struct {
}

func (s *ArticleService) Add(a Article) error {
	return nil
}
