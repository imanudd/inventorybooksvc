package registry

import (
	"github.com/imanudd/inventorybooksvc/internal/core/port/inbound/service"
)

type ServiceRegistry interface {
	GetAuthService() service.AuthService
	GetBookService() service.BookService
	GetAuthorService() service.AuthorService
}
