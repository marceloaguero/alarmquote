package articles

const (
	// ErrNotFound means that the article could not be found in the reposutory
	ErrNotFound = RepoErr("Could not found the article")

	// ErrAlreadyExists means that the article exists in the repository
	ErrAlreadyExists = RepoErr("Article already exists")

	ErrNoCode = RepoErr("Article has no code")

	ErrNoName = RepoErr("Article has no name")

	ErrNoCategory = RepoErr("Article has no category")
)

// RepoErr are errors that can happen when using the articles repository
type RepoErr string

func (e RepoErr) Error() string {
	return string(e)
}

// Article represents a single article
type Article struct {
	Code        string
	Name        string
	Description string
	Category    string
	Price       float64
}

// Articles is the mocked storage for articles
type Articles map[string]Article

// Add saves a new article in permanent storage
func (repo Articles) Add(a Article) error {
	if a.Code == "" {
		return ErrNoCode
	}

	if a.Name == "" {
		return ErrNoName
	}

	if a.Category == "" {
		return ErrNoCategory
	}

	_, err := repo.GetByCode(a.Code)

	switch err {
	case ErrNotFound:
		repo[a.Code] = a
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}

// GetByCode retrieves an article searching by its code
func (repo Articles) GetByCode(code string) (Article, error) {
	article, ok := repo[code]
	if !ok {
		return Article{}, ErrNotFound
	}
	return article, nil
}
