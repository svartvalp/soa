package main

import (
	"context"

	log "github.com/svartvalp/soa/service/logger"

	"github.com/soa/search-api/internal/config"
	"github.com/soa/search-api/internal/controllers/search"
	search_service "github.com/soa/search-api/internal/pkg/search-service"
	"github.com/svartvalp/soa/service/server"
)

// @title    Search API
// @version  1.0
// @host     localhost:7001
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	logger := log.LoggerFromContext(ctx)
	// Config
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		logger.Fatal(err)
	}

	// Services
	searchService, err := search_service.NewService()
	if err != nil {
		logger.Fatal(err)
	}

	// Controllers
	productController := search.NewController(searchService)

	srv := server.NewServer(
		&server.Config{
			Host: cfg.Server.Host,
			Port: cfg.Server.Port,
			Controllers: []server.Controller{
				productController,
			},
		},
		server.WithLogger,
		server.WithMetrics,
	)
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
