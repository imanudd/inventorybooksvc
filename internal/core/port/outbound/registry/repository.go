package registry

import "github.com/imanudd/inventorybooksvc/internal/core/port/outbound/repository"

type RepositoryRegistry interface {
	GetBookRepository() repository.BookRepository
	GetAuthorRepository() repository.AuthorRepository
	GetUserRepository() repository.UserRepository
}
