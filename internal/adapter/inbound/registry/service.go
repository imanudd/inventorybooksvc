package registry

import (
	"github.com/imanudd/inventorybooksvc/config"
	inregistry "github.com/imanudd/inventorybooksvc/internal/core/port/inbound/registry"
	inservice "github.com/imanudd/inventorybooksvc/internal/core/port/inbound/service"
	outregistry "github.com/imanudd/inventorybooksvc/internal/core/port/outbound/registry"
	"github.com/imanudd/inventorybooksvc/internal/core/service"
)

type ServiceRegistryConfig struct {
	Config     *config.MainConfig
	Repository outregistry.RepositoryRegistry
}

type ServiceRegistry struct {
	bookService   inservice.BookService
	authorService inservice.AuthorService
	authService   inservice.AuthService
}

func NewServiceRegistry(registry *ServiceRegistryConfig) inregistry.ServiceRegistry {
	return &ServiceRegistry{
		bookService:   service.NewBookService(registry.Config, registry.Repository),
		authorService: service.NewAuthorService(registry.Config, registry.Repository),
		authService:   service.NewAuthService(registry.Config, registry.Repository),
	}
}

func (s *ServiceRegistry) GetAuthService() inservice.AuthService {
	return s.authService
}

func (s *ServiceRegistry) GetBookService() inservice.BookService {
	return s.bookService
}

func (s *ServiceRegistry) GetAuthorService() inservice.AuthorService {
	return s.authorService
}
