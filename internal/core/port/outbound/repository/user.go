package repository

import (
	"context"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetByUsernameOrEmail(ctx context.Context, req *domain.GetByUsernameOrEmail) (*domain.User, error)
	RegisterUser(ctx context.Context, req *domain.User) error
}
