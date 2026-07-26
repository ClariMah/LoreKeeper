package ports

import (
	"context"
	"lorekeeper/internal/domain"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *domain.Book) error
	GetById(ctx context.Context, id int32) (*domain.Book, error)
	ListBooks(ctx context.Context) ([]*domain.Book, error)
	UpdateBook(ctx context.Context, book *domain.Book) error
	DeleteBook(ctx context.Context, id int32) error
}
