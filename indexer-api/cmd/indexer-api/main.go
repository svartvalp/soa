package main

import (
	"context"
	"log"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/kafka"
	product_service "github.com/soa/indexer-api/internal/pkg/product-service"
	search_service "github.com/soa/indexer-api/internal/pkg/search-service"
	product_api "github.com/soa/indexer-api/internal/requester"
)

func main() {
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {

	}

	requester := product_api.NewRequester()

	// Services
	productService := product_service.NewService(cfg, requester)
	searchService := search_service.NewService(cfg, requester)

	cons := kafka.NewConsumer(cfg, productService, searchService)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Indexer run")
	cons.Start(context.Background())
}
