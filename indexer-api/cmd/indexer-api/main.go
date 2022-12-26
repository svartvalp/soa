package main

import (
	"context"
	"log"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/db"
	"github.com/soa/indexer-api/internal/db/indexer"
	"github.com/soa/indexer-api/internal/kafka"
	product_service "github.com/soa/indexer-api/internal/pkg/product-service"
	"github.com/soa/indexer-api/internal/requester"
	product_api "github.com/soa/indexer-api/internal/requester/product-api"
	search_api "github.com/soa/indexer-api/internal/requester/search-api"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// API
	requester := requester.NewRequester()
	productAPI := product_api.New(cfg, requester)
	searchAPI := search_api.New(cfg, requester)

	// Repository
	conn, err := db.NewWrapper(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	repository := indexer.NewRepository(conn)

	// Service
	productService := product_service.New(repository, productAPI, searchAPI)

	cons := kafka.NewConsumer(cfg, productService)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Indexer run")
	cons.Start(context.Background())
}
