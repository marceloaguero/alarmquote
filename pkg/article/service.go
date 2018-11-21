package article

import "github.com/marceloaguero/alarmquote"

// Ensure ArticleService implements alarmquote.ArticleService
var _ alarmquote.ArticleService = &ArticleService{}
