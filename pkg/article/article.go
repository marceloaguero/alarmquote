package article

/*
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
*/
