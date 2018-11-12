package articles

import (
	"testing"
)

func TestArticlesAdd(t *testing.T) {
	article := Article{Code: "P1101216",
		Name:     "N4-MPXH",
		Category: "Centrales",
		Price:    2168.00,
	}

	t.Run("New article", func(t *testing.T) {
		repo := Articles{}
		err := repo.Add(article)
		if err != nil {
			t.Errorf("repo.Add should not return an error, got: %v", err)
		}

		// Retrieve article to verify if it was saved
		article, err = repo.GetByCode("P1101216")
		if err != nil {
			t.Errorf("GetByCode should not retrieve an error, got: %v", err)
		}
		if article.Code != "P1101216" {
			t.Errorf("GetByCode, want: '%s', got: '%s'", "P1101216", article.Code)
		}
	})

	t.Run("Existing article", func(t *testing.T) {
		repo := Articles{}
		a := Article{Code: "P1101216"}
		repo["P1101216"] = a
		err := repo.Add(a)
		if err == nil {
			t.Errorf("Adding an existing article should retrieve an error")
		}
	})
}
