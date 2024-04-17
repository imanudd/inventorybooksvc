package cmd

import (
	"log"

	"github.com/imanudd/inventorybooksvc/config"
	inregistry "github.com/imanudd/inventorybooksvc/internal/adapter/inbound/registry"
	"github.com/imanudd/inventorybooksvc/internal/adapter/inbound/rest"
	outregistry "github.com/imanudd/inventorybooksvc/internal/adapter/outbound/registry"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "rest",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		pgDB := InitPostgreSQL(cfg)

		if cfg.LogMode {
			pgDB = pgDB.Debug()
		}

		repoRegistry := outregistry.NewRepositoryRegistry(pgDB)

		serviceRegistry := inregistry.NewServiceRegistry(&inregistry.ServiceRegistryConfig{
			Config:     cfg,
			Repository: repoRegistry,
		})

		rest := rest.New(cfg)

		if err := rest.RegisterHandler(cfg, repoRegistry, serviceRegistry); err != nil {
			return
		}

		if err := rest.Serve(); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}

	},
}
