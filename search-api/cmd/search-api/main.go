package main

import (
	"context"
	"log"

	"github.com/soa/search-api/internal/config"
	"github.com/soa/search-api/internal/controllers/search"
	"github.com/soa/search-api/internal/db"
	search_repo "github.com/soa/search-api/internal/db/search"
	search_service "github.com/soa/search-api/internal/pkg/search-service"
	"github.com/soa/search-api/internal/server"
)

func main() {
	ctx := context.Background()

	// Config
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.NewWrapper(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// repository
	searchRepository := search_repo.NewRepository(conn)

	// Services
	searchService := search_service.NewService(searchRepository)

	// Controllers
	productController := search.NewController(searchService)

	srv := server.NewServer(
		cfg,
		productController,
	)

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
