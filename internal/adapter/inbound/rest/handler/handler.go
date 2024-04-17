package handler

import (
	"github.com/imanudd/inventorybooksvc/internal/core/port/inbound/registry"
)

type Handler struct {
	service registry.ServiceRegistry
}

func NewHandler(service registry.ServiceRegistry) *Handler {
	return &Handler{
		service: service,
	}
}
