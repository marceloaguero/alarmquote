package article

import (
	"testing"

	"github.com/marceloaguero/alarmquote"
)

func TestAdd(t *testing.T) {
	repo := newMockRepo()
	s := NewService(repo)

	var articlesTests = []struct {
		desc    string
		article alarmquote.Article
		want    error
	}{
		{desc: "Valid article",
			article: alarmquote.Article{
				ID:       "P1101216",
				Name:     "N4-MPXH",
				Category: "Centrales",
				Price:    2168.00},
			want: nil,
		},
		{desc: "Article without ID",
			article: alarmquote.Article{
				Name:     "N4-MPXH",
				Category: "Centrales",
				Price:    2168.00},
			want: alarmquote.ErrArticleIDRequired,
		},
		{desc: "Article without name",
			article: alarmquote.Article{
				ID:       "P1101216",
				Category: "Centrales",
				Price:    2168.00},
			want: alarmquote.ErrArticleNameRequired,
		},
		{desc: "Article without category",
			article: alarmquote.Article{
				ID:    "P1101216",
				Name:  "N4-MPXH",
				Price: 2168.00},
			want: alarmquote.ErrArticleCategoryRequired,
		},
	}

	for _, tt := range articlesTests {
		t.Run(tt.desc, func(t *testing.T) {
			err := s.Add(tt.article)
			if err != tt.want {
				t.Errorf("got: %v, want: %v", err, tt.want)
			}
		})
	}

	t.Run("Existing article", func(t *testing.T) {
		repo := newMockRepo()
		s := NewService(repo)

		a := alarmquote.Article{
			ID:       "P1101216",
			Name:     "N4-MPXH",
			Category: "Centrales",
			Price:    2168.00,
		}

		// Add the article. This should be OK
		err := s.Add(a)
		if err != nil {
			t.Errorf("Adding a valid new article should NOT retrieve an error")
		}

		// Add the article again. This shouldn't be allowed
		err = s.Add(a)
		want := alarmquote.ErrArticleExists
		if err != want {
			t.Errorf("got: %v, want: %v", err, want)
		}
	})

}

type mockRepo struct {
	articles map[alarmquote.ArticleID]alarmquote.Article
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		articles: map[alarmquote.ArticleID]alarmquote.Article{},
	}
}

func (r *mockRepo) Insert(a alarmquote.Article) error {
	r.articles[a.ID] = a
	return nil
}

func (r *mockRepo) Retrieve(id alarmquote.ArticleID) (*alarmquote.Article, error) {
	article, ok := r.articles[id]
	if !ok {
		return nil, alarmquote.ErrArticleNotFound
	}

	return &article, nil
}
