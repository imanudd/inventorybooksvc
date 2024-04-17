package registry

import (
	adapter "github.com/imanudd/inventorybooksvc/internal/adapter/outbound/repository"
	"github.com/imanudd/inventorybooksvc/internal/core/port/outbound/registry"
	port "github.com/imanudd/inventorybooksvc/internal/core/port/outbound/repository"
	"gorm.io/gorm"
)

type RepositoryRegistry struct {
	db *gorm.DB
}

func NewRepositoryRegistry(db *gorm.DB) registry.RepositoryRegistry {
	return &RepositoryRegistry{
		db: db,
	}
}

func (r *RepositoryRegistry) GetBookRepository() port.BookRepository {
	return adapter.NewBookRepository(r.db)
}

func (r *RepositoryRegistry) GetAuthorRepository() port.AuthorRepository {
	return adapter.NewAuthorRepository(r.db)
}

func (r *RepositoryRegistry) GetUserRepository() port.UserRepository {
	return adapter.NewUserRepository(r.db)
}
