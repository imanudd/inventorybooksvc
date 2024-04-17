package service

import (
	"context"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

type BookService interface {
	GetDetailBook(ctx context.Context, id int) (*domain.DetailBook, error)
	DeleteBook(ctx context.Context, id int) error
	UpdateBook(ctx context.Context, req *domain.UpdateBookRequest) error
	AddBook(ctx context.Context, req *domain.CreateBookRequest) error
}
