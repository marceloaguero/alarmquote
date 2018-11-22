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

func TestEdit(t *testing.T) {
	repo := newMockRepo()
	s := NewService(repo)

	a := alarmquote.Article{
		ID:       "P1101216",
		Name:     "N4-MPXH",
		Category: "Centrales",
		Price:    2168.00,
	}

	// Add a new article. This should be OK
	err := s.Add(a)
	if err != nil {
		t.Errorf("Adding a valid new article should NOT retrieve an error")
	}

	// Test cases
	var articlesTests = []struct {
		desc    string
		id      alarmquote.ArticleID
		article alarmquote.Article
		want    error
	}{
		{desc: "Valid edition",
			id: "P1101216",
			article: alarmquote.Article{
				ID:       "P1101216",
				Name:     "N4-MPXH",
				Category: "Centrales",
				Price:    2000.00},
			want: nil,
		},
		{desc: "Article without ID",
			id: "P1101216",
			article: alarmquote.Article{
				Name:     "N4-MPXH",
				Category: "Centrales",
				Price:    2168.00},
			want: alarmquote.ErrArticleIDRequired,
		},
		{desc: "Article without name",
			id: "P1101216",
			article: alarmquote.Article{
				ID:       "P1101216",
				Category: "Centrales",
				Price:    2168.00},
			want: alarmquote.ErrArticleNameRequired,
		},
		{desc: "Article without category",
			id: "P1101216",
			article: alarmquote.Article{
				ID:    "P1101216",
				Name:  "N4-MPXH",
				Price: 2168.00},
			want: alarmquote.ErrArticleCategoryRequired,
		},
	}

	for _, tt := range articlesTests {
		t.Run(tt.desc, func(t *testing.T) {
			err := s.Edit(tt.id, tt.article)
			if err != tt.want {
				t.Errorf("got: %v, want: %v", err, tt.want)
			}
		})
	}

	// Test a valid edition (price)
	a = alarmquote.Article{
		ID:       "P1101216",
		Name:     "N4-MPXH",
		Category: "Centrales",
		Price:    3000.00,
	}

	t.Run("Edit price", func(t *testing.T) {
		err := s.Edit(a.ID, a)
		if err != nil {
			t.Errorf("Editing price should NOT retrieve an error")
		}

		edited, err := s.GetByID(a.ID)
		if a.Price != edited.Price {
			t.Errorf("got: %f, want: %f", edited.Price, a.Price)
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

func (r *mockRepo) Retrieve(id alarmquote.ArticleID) (*alarmquote.Article, error) {
	article, ok := r.articles[id]
	if !ok {
		return nil, alarmquote.ErrArticleNotFound
	}

	return &article, nil
}

func (r *mockRepo) Insert(a alarmquote.Article) error {
	r.articles[a.ID] = a
	return nil
}

func (r *mockRepo) Modify(id alarmquote.ArticleID, a alarmquote.Article) error {
	r.articles[a.ID] = a
	return nil
}
