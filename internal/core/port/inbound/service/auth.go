package service

import (
	"context"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error)
	Register(ctx context.Context, req *domain.RegisterRequest) error
}
