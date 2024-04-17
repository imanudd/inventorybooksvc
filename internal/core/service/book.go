package service

import (
	"context"
	"errors"
	"time"

	"github.com/imanudd/inventorybooksvc/config"
	"github.com/imanudd/inventorybooksvc/internal/core/domain"
	"github.com/imanudd/inventorybooksvc/internal/core/port/inbound/service"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/registry"
)

type bookService struct {
	config     config.MainConfig
	repository registry.RepositoryRegistry
}

func NewBookService(config *config.MainConfig, repository registry.RepositoryRegistry) service.BookService {
	return &bookService{
		config:     *config,
		repository: repository,
	}
}

func (s *bookService) DeleteBook(ctx context.Context, id int) error {
	book, err := s.repository.GetBookRepository().GetByID(ctx, id)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	return s.repository.GetBookRepository().Delete(ctx, id)
}

func (s *bookService) GetDetailBook(ctx context.Context, id int) (*domain.DetailBook, error) {
	book, err := s.repository.GetBookRepository().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	author, err := s.repository.GetAuthorRepository().GetByID(ctx, book.AuthorID)
	if err != nil {
		return nil, err
	}

	resp := &domain.DetailBook{
		ID:         book.ID,
		AuthorID:   book.AuthorID,
		AuthorName: author.Name,
		BookName:   book.BookName,
		Title:      book.Title,
		Price:      book.Price,
		CreatedAt:  book.CreatedAt,
	}

	return resp, nil
}

func (s *bookService) UpdateBook(ctx context.Context, req *domain.UpdateBookRequest) error {
	book, err := s.repository.GetBookRepository().GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	author, err := s.repository.GetAuthorRepository().GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return s.repository.GetBookRepository().Update(ctx, &domain.Book{
		ID:       req.ID,
		AuthorID: req.AuthorID,
		BookName: req.BookName,
		Title:    req.Title,
		Price:    req.Price,
	})
}

func (s *bookService) AddBook(ctx context.Context, req *domain.CreateBookRequest) error {
	author, err := s.repository.GetAuthorRepository().GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	book := &domain.Book{
		AuthorID:  req.AuthorID,
		BookName:  req.BookName,
		Title:     req.Title,
		Price:     req.Price,
		CreatedAt: time.Now(),
	}

	return s.repository.GetBookRepository().Create(ctx, book)
}
