package articles

// Article represents a single article
type Article struct {
	Code        string
	Name        string
	Description string
	Category    string
	Price       float64
}

// Articles is the mocked storage for articles
type Articles []Article

// Save saves an article in permanent storage
func (a *Article) Save() error {
	return nil
}

func main() {

}
