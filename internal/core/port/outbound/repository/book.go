package repository

import (
	"context"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

type BookRepository interface {
	GetListBookByAuthorID(ctx context.Context, authorID int) ([]*domain.Book, error)
	DeleteBookByAuthorID(ctx context.Context, authorID, bookID int) error
	GetByID(ctx context.Context, id int) (*domain.Book, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, req *domain.Book) error
	Create(ctx context.Context, req *domain.Book) error
}
