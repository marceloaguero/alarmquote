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

func TestEdit(t *testing.T) {
	repo := newMockRepo()
	s := NewService(repo)

	c := alarmquote.Category{Name: "Centrales"}

	// Add a new category. This should be OK.
	err := s.Add(c)
	if err != nil {
		t.Errorf("Adding a valid new category should NOT retrieve an error")
	}

	// Test cases.
	var categoriesTests = []struct {
		desc     string
		n        alarmquote.CategoryName
		category alarmquote.Category
		want     error
	}{
		{desc: "Valid edition",
			n:        "Centrales",
			category: alarmquote.Category{Name: "Centrales MPX"},
			want:     nil,
		},
		{desc: "Category without name",
			n:        "Centrales",
			category: alarmquote.Category{},
			want:     alarmquote.ErrCategoryNameRequired,
		},
	}

	for _, tt := range categoriesTests {
		t.Run(tt.desc, func(t *testing.T) {
			err := s.Edit(tt.n, tt.category)
			if err != tt.want {
				t.Errorf("got: %v, want: %v", err, tt.want)
			}
		})
	}

	// Test a valid category edition.
	repo = newMockRepo()
	s = NewService(repo)

	c = alarmquote.Category{Name: "Centrales"}

	// Add a new category. This should be OK.
	err = s.Add(c)
	if err != nil {
		t.Errorf("Adding a valid new category should NOT retrieve an error")
	}

	cat := alarmquote.Category{Name: "Centrales"}

	t.Run("Edit category", func(t *testing.T) {
		err := s.Edit(c.Name, cat)
		if err != nil {
			t.Errorf("Editing category should NOT retrieve an error")
		}

		edited, err := s.Category(cat.Name)
		if edited.Name != cat.Name {
			t.Errorf("got: %s, want: %s", edited.Name, cat.Name)
		}
	})
}

func TestDelete(t *testing.T) {
	repo := newMockRepo()
	s := NewService(repo)

	c := alarmquote.Category{Name: "Centrales"}

	// Add a new category. This should be OK.
	err := s.Add(c)
	if err != nil {
		t.Errorf("Adding a valid new category should NOT retrieve an error")
	}

	// Test cases.

	// Non existent category. Deletion should fail
	var nonExistentName alarmquote.CategoryName = "Sensores"
	err = s.Delete(nonExistentName)
	want := alarmquote.ErrCategoryNotFound
	if err != want {
		t.Errorf("Deleting non existent Category, got: %v, want %v", err, want)
	}

	// Existent category. Deletion should success
	var existentName alarmquote.CategoryName = "Centrales"
	err = s.Delete(existentName)
	want = nil
	if err != want {
		t.Errorf("Deleting existent Category, got: %v, want %v", err, want)
	}
	// Retrieving the deleted category should give an error
	_, err = s.Category(existentName)
	want = alarmquote.ErrCategoryNotFound
	if err != want {
		t.Errorf("Get a deleted category, got: %v, want: %v", err, want)
	}
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

func (r *mockRepo) Modify(n alarmquote.CategoryName, c alarmquote.Category) error {
	for _, cat := range r.categories {
		if cat.Name == n {
			cat = c
		}
	}

	return nil
}

func (r *mockRepo) Delete(n alarmquote.CategoryName) error {
	for i, c := range r.categories {
		if c.Name == n {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			break
		}
	}
	return nil
}
