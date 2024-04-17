package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/repository"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Create(ctx context.Context, req *domain.Author) error {
	return r.db.WithContext(ctx).Model(&domain.Author{}).Create(&req).Error
}

func (r *AuthorRepository) GetByName(ctx context.Context, name string) (*domain.Author, error) {
	var author domain.Author
	db := r.db.Model(&domain.Author{}).Where("name ilike ?", name).First(&author)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) GetByID(ctx context.Context, id int) (*domain.Author, error) {
	var author domain.Author
	db := r.db.Model(&domain.Author{}).Where("id = ?", id).First(&author)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return &author, nil
}
