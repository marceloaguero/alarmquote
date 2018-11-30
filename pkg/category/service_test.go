package category

import (
	"testing"

	"github.com/marceloaguero/alarmquote"
)

func TestAdd(t *testing.T) {
	repo := newMockRepo()
	s := NewService(repo)

	var categoriesTests = []struct {
		desc     string
		category alarmquote.Category
		want     error
	}{
		{desc: "Valid category",
			category: alarmquote.Category{Name: "Centrales"},
			want:     nil,
		},
		{desc: "Category without name",
			category: alarmquote.Category{Name: ""},
			want:     alarmquote.ErrCategoryNameRequired,
		},
	}

	for _, tt := range categoriesTests {
		t.Run(tt.desc, func(t *testing.T) {
			err := s.Add(tt.category)
			if err != tt.want {
				t.Errorf("got: %v, want: %v", err, tt.want)
			}
		})
	}

	t.Run("Existing category", func(t *testing.T) {
		repo := newMockRepo()
		s := NewService(repo)

		c := alarmquote.Category{Name: "Centrales"}

		// Add the category. This should be OK.
		err := s.Add(c)
		if err != nil {
			t.Errorf("Adding a valid new category should NOT retrieve an error")
		}

		// Add the category again. This shouldn't be allowed.
		err = s.Add(c)
		want := alarmquote.ErrCategoryExists
		if err != want {
			t.Errorf("got: %v, want: %v", err, want)
		}
	})
}

type mockRepo struct {
	categories []alarmquote.Category
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		categories: []alarmquote.Category{},
	}
}

func (r *mockRepo) Retrieve(n alarmquote.CategoryName) (*alarmquote.Category, error) {
	for _, c := range r.categories {
		if c.Name == n {
			return &alarmquote.Category{Name: c.Name}, nil
		}
	}

	return nil, alarmquote.ErrCategoryNotFound
}

func (r *mockRepo) Insert(c alarmquote.Category) error {
	r.categories = append(r.categories, c)
	return nil
}
