package domain

type Book struct {
	ID           int32  `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Publisher    string `json:"publisher"`
	ReleasedYear int32  `json:"released_year"`
	PageNumber   int32  `json:"page_number"`
	Category     string `json:"category"`
	Language     string `json:"language"`
	ISBN         string `json:"isbn"`
}

func NewBook(title, author, publisher string, releasedyear, pagenumber int32, category, language, isbn string) (*Book, error) {
	if title == "" {
		return nil, ErrTitleRequired
	}
	if author == "" {
		return nil, ErrAuthorRequired
	}
	if pagenumber <= 0 {
		return nil, ErrInvalidPages
	}
	return &Book{
		Title:        title,
		Author:       author,
		Publisher:    publisher,
		ReleasedYear: releasedyear,
		PageNumber:   pagenumber,
		Category:     category,
		Language:     language,
		ISBN:         isbn,
	}, nil
}
