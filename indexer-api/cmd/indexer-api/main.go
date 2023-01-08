package main

import (
	"context"

	_ "github.com/soa/indexer-api/docs"
	product_api "github.com/soa/indexer-api/internal/clients/product"
	search_api "github.com/soa/indexer-api/internal/clients/search"
	"github.com/soa/indexer-api/internal/config"
	indexer_controller "github.com/soa/indexer-api/internal/controllers/indexer"
	"github.com/soa/indexer-api/internal/db"
	"github.com/soa/indexer-api/internal/db/indexer"
	"github.com/soa/indexer-api/internal/kafka"
	product_service "github.com/soa/indexer-api/internal/pkg/product-service"
	"github.com/svartvalp/soa/service/logger"
	"github.com/svartvalp/soa/service/server"
)

// @title   Indexer API
// @version 1.0

// @host     localhost:7004
// @BasePath /api/v1
func main() {
	ctx := logger.ContextWithLogger(context.Background())
	log := logger.LoggerFromContext(ctx)

	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// API
	productClient := product_api.New(cfg)
	searchClient := search_api.New(cfg)

	// Repository
	conn, err := db.NewWrapper(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	repository := indexer.NewRepository(conn)

	// Service
	productService := product_service.New(repository, productClient, searchClient)

	// Controller
	indexerController := indexer_controller.New(productService)

	// Kafka
	cons := kafka.NewConsumer(cfg, productService)
	if err != nil {
		// log.Fatal(err)
	}

	go func() {
		cons.Start(ctx)
	}()

	srv := server.NewServer(&server.Config{
		Host:        cfg.Server.Host,
		Port:        cfg.Server.Port,
		Controllers: []server.Controller{indexerController},
	},
		server.WithLogger,
	)
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
