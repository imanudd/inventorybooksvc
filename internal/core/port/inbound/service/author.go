package service

import (
	"context"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

type AuthorService interface {
	DeleteBookByAuthor(ctx context.Context, id, bookId int) error
	GetListBookByAuthor(ctx context.Context, id int) ([]*domain.Book, error)
	AddAuthorBook(ctx context.Context, req *domain.AddAuthorBookRequest) error
	CreateAuthor(ctx context.Context, req *domain.CreateAuthorRequest) error
}
