package service

import (
	"context"
	"errors"

	"github.com/imanudd/inventorybooksvc/config"
	"github.com/imanudd/inventorybooksvc/internal/core/domain"
	"github.com/imanudd/inventorybooksvc/internal/core/port/inbound/service"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/registry"
)

type authorService struct {
	config     config.MainConfig
	repository registry.RepositoryRegistry
}

func (a *authorService) AddAuthorBook(ctx context.Context, req *domain.AddAuthorBookRequest) error {
	author, err := a.repository.GetAuthorRepository().GetByID(ctx, req.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return a.repository.GetBookRepository().Create(ctx, &domain.Book{
		AuthorID: author.ID,
		BookName: req.BookName,
		Title:    req.Title,
		Price:    req.Price,
	})
}

func (a *authorService) CreateAuthor(ctx context.Context, req *domain.CreateAuthorRequest) error {
	author, err := a.repository.GetAuthorRepository().GetByName(ctx, req.Name)
	if err != nil {
		return err
	}

	if author != nil {
		return errors.New("author already exist")
	}

	return a.repository.GetAuthorRepository().Create(ctx, &domain.Author{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})
}

func (a *authorService) DeleteBookByAuthor(ctx context.Context, id int, bookId int) error {
	book, err := a.repository.GetBookRepository().GetByID(ctx, bookId)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	author, err := a.repository.GetAuthorRepository().GetByID(ctx, id)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("author not found")
	}

	return a.repository.GetBookRepository().DeleteBookByAuthorID(ctx, author.ID, book.ID)
}

func (a *authorService) GetListBookByAuthor(ctx context.Context, id int) ([]*domain.Book, error) {
	author, err := a.repository.GetAuthorRepository().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, errors.New("author not found")
	}

	books, err := a.repository.GetBookRepository().GetListBookByAuthorID(ctx, author.ID)
	if err != nil {
		return nil, err
	}

	if len(books) < 1 {
		return nil, errors.New("book not found")
	}

	return books, nil
}

func NewAuthorService(config *config.MainConfig, repository registry.RepositoryRegistry) service.AuthorService {
	return &authorService{
		config:     *config,
		repository: repository,
	}
}
