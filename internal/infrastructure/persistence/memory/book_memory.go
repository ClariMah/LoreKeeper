package memory

import (
	"context"
	"lorekeeper/internal/domain"
)

type BookMemoryRepository struct {
	Data []*domain.Book
}

func NewBooksMemoryRepository() *BookMemoryRepository {
	return &BookMemoryRepository{
		Data: []*domain.Book{},
	}
}

func (rb *BookMemoryRepository) CreateBook(ctx context.Context, book *domain.Book) error {
	rb.Data = append(rb.Data, book)
	return nil
}

func (rb *BookMemoryRepository) GetById(ctx context.Context, id int32) (*domain.Book, error) {
	for _, book := range rb.Data {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, domain.ErrBookNotFound
}

func (rb *BookMemoryRepository) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	return rb.Data, nil
}

func (rb *BookMemoryRepository) UpdateBook(ctx context.Context, book *domain.Book) error {
	for i, b := range rb.Data {
		if b.ID == book.ID {
			rb.Data[i] = book
			return nil
		}
	}
	return domain.ErrBookNotFound
}

func (rb *BookMemoryRepository) DeleteBook(ctx context.Context, id int32) error {
	for i, b := range rb.Data {
		if b.ID == id {
			rb.Data = append(rb.Data[:i], rb.Data[i+1:]...)
			return nil
		}
	}
	return domain.ErrBookNotFound
}
