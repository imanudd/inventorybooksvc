package repository

import (
	"context"
	"errors"

	"github.com/imanudd/inventorybooksvc/internal/core/domain"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/repository"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) GetListBookByAuthorID(ctx context.Context, authorID int) ([]*domain.Book, error) {
	var books []*domain.Book

	db := r.db.WithContext(ctx).Model(&domain.Book{}).Where("author_id = ?", authorID).Find(&books)
	if err := db.Error; err != nil {
		return nil, err
	}

	return books, nil

}

func (r *BookRepository) DeleteBookByAuthorID(ctx context.Context, authorID int, bookID int) error {
	return r.db.WithContext(ctx).Model(&domain.Book{}).Delete("id = ? and author_id = ?", bookID, authorID).Error
}

func (r *BookRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&domain.Book{}).Delete("id = ?", id).Error
}

func (r *BookRepository) GetByID(ctx context.Context, id int) (*domain.Book, error) {
	var book domain.Book
	db := r.db.Model(&domain.Book{}).Where("id = ?", id).First(&book)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err := db.Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) Update(ctx context.Context, req *domain.Book) error {
	return r.db.WithContext(ctx).Omit("id").Model(&domain.Book{}).Where("id = ?", req.ID).Updates(&req).Error
}

func (r *BookRepository) Create(ctx context.Context, req *domain.Book) error {
	db := r.db.Model(&domain.Book{}).Create(&req)

	return db.Error
}
