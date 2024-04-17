package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/imanudd/inventorybooksvc/config"
	"github.com/imanudd/inventorybooksvc/internal/core/domain"
	"github.com/imanudd/inventorybooksvc/internal/core/port/inbound/service"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/registry"
	"github.com/imanudd/inventorybooksvc/pkg/auth"
)

type authService struct {
	config     config.MainConfig
	repository registry.RepositoryRegistry
}

func NewAuthService(config *config.MainConfig, repository registry.RepositoryRegistry) service.AuthService {
	return &authService{
		config:     *config,
		repository: repository,
	}
}

func (a *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := a.repository.GetUserRepository().GetByUsernameOrEmail(ctx, &domain.GetByUsernameOrEmail{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user is not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	auth := auth.NewAuth(&a.config)
	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Username: req.Username,
		Token:    token,
	}, nil
}

func (a *authService) Register(ctx context.Context, req *domain.RegisterRequest) (err error) {

	user, err := a.repository.GetUserRepository().GetByUsernameOrEmail(ctx, &domain.GetByUsernameOrEmail{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return
	}

	if user != nil {
		return errors.New("user is already exist")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error when hashing password")
	}

	return a.repository.GetUserRepository().RegisterUser(ctx, &domain.User{
		Username: req.Username,
		Password: string(hash),
		Email:    req.Email,
	})

}
