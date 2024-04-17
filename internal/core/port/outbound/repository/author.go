package repository

import (
	"context"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

type AuthorRepository interface {
	Create(ctx context.Context, req *domain.Author) error
	GetByName(ctx context.Context, name string) (*domain.Author, error)
	GetByID(ctx context.Context, id int) (*domain.Author, error)
}
