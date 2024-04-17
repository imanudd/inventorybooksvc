package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	db := r.db.Model(&user).Where("id = ?", id).First(&user)
	if err := db.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByUsernameOrEmail(ctx context.Context, req *domain.GetByUsernameOrEmail) (*domain.User, error) {
	var user *domain.User
	db := r.db.WithContext(ctx).Model(&domain.User{}).Where("username ilike ? or email ilike ?", req.Username, req.Email).First(&user)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) RegisterUser(ctx context.Context, req *domain.User) error {
	db := r.db.WithContext(ctx).Model(&domain.User{}).Create(&req)

	return db.Error
}
