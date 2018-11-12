package articles

import (
	"testing"
)

func TestArticlesNew(t *testing.T) {
	article := Article{Code: "P1101216",
		Name:     "N4-MPXH",
		Category: "Centrales",
		Price:    2168.00,
	}

	err := article.Save()
	if err != nil {
		t.Errorf("Could not save article, want err == nil, got err: %v", err)
	}
}
