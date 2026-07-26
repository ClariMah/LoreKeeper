package application

import (
	"context"
	"lorekeeper/internal/domain"
	"lorekeeper/internal/ports"
)

type BookUseCase struct {
	repository ports.BookRepository
}

func NewUseCase(repository ports.BookRepository) *BookUseCase {
	return &BookUseCase{
		repository: repository,
	}
}

func (uc *BookUseCase) CreateBook(ctx context.Context, title, author, publisher string, releasedyear, pagenumber int32, category, language, isbn string) (*domain.Book, error) {
	book, err := domain.NewBook(title, author, publisher, releasedyear, pagenumber, category, language, isbn)
	if err != nil {
		return nil, err
	}
	if err := uc.repository.CreateBook(ctx, book); err != nil {
		return nil, err
	}
	return book, nil
}

func (uc *BookUseCase) GetBookById(ctx context.Context, id int32) (*domain.Book, error) {
	return uc.repository.GetById(ctx, id)
}

func (uc *BookUseCase) ListBooks(ctx context.Context) ([]*domain.Book, error) {
	return uc.repository.ListBooks(ctx)
}

func (uc *BookUseCase) UpdateBook(ctx context.Context, book *domain.Book) error {
	return uc.repository.UpdateBook(ctx, book)
}

func (uc *BookUseCase) DeleteBook(ctx context.Context, id int32) error {
	return uc.repository.DeleteBook(ctx, id)
}
