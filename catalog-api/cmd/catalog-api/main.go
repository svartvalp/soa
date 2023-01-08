package main

import (
	"context"

	product_client "github.com/soa/catalog-api/internal/clients/product"
	search_client "github.com/soa/catalog-api/internal/clients/search"
	"github.com/soa/catalog-api/internal/config"
	catalog_controller "github.com/soa/catalog-api/internal/controllers/catalog"
	"github.com/soa/catalog-api/internal/pkg/catalog"
	"github.com/soa/catalog-api/internal/s3"
	"github.com/svartvalp/soa/service/logger"
	"github.com/svartvalp/soa/service/server"
)

// @title   Catalog API
// @version 1.0

// @host     localhost:7003
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	log := logger.LoggerFromContext(ctx)

	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// APIs
	productAPI := product_client.New(cfg)
	searchAPI := search_client.New(cfg)

	s3Cl, err := s3.NewS3(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Services
	catalogService := catalog.New(productAPI, searchAPI, s3Cl)

	// Controllers
	catalogController := catalog_controller.New(catalogService)

	srv := server.NewServer(
		&server.Config{
			Host:        cfg.Server.Host,
			Port:        cfg.Server.Port,
			Controllers: []server.Controller{catalogController},
		},
		server.WithLogger,
		server.WithMetrics,
	)

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
