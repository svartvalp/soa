package main

import (
	"log"

	"github.com/soa/search-api/internal/config"
	"github.com/soa/search-api/internal/controllers/search"
	search_service "github.com/soa/search-api/internal/pkg/search-service"
	"github.com/soa/search-api/internal/server"
)

func main() {
	// Config
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Services
	searchService, err := search_service.NewService()
	if err != nil {
		panic(err)
	}

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
