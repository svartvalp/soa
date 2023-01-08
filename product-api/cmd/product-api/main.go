package main

import (
	"context"

	_ "github.com/soa/product-api/docs"
	"github.com/soa/product-api/internal/config"
	"github.com/soa/product-api/internal/controllers/category"
	"github.com/soa/product-api/internal/controllers/characteristic"
	"github.com/soa/product-api/internal/controllers/product"
	"github.com/soa/product-api/internal/db"
	category_repo "github.com/soa/product-api/internal/db/category"
	characteristic_repo "github.com/soa/product-api/internal/db/characteristic"
	product_repo "github.com/soa/product-api/internal/db/product"
	"github.com/soa/product-api/internal/kafka"
	category_service "github.com/soa/product-api/internal/pkg/category-service"
	characteristic_service "github.com/soa/product-api/internal/pkg/characteristic-service"
	product_service "github.com/soa/product-api/internal/pkg/product-service"
	"github.com/soa/product-api/internal/s3"
	"github.com/svartvalp/soa/service/logger"
	"github.com/svartvalp/soa/service/server"
)

// @title   Product API
// @version 1.0

// @host     localhost:7002
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	log := logger.LoggerFromContext(ctx)

	// Config
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.NewWrapper(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	s3Client, err := s3.NewS3(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Kafka
	producer := kafka.NewProducer(cfg)

	// repository
	productRepository := product_repo.NewRepository(conn)
	categoryRepository := category_repo.NewRepository(conn)
	characteristicRepository := characteristic_repo.NewRepository(conn)

	// Services
	categoryService := category_service.NewService(categoryRepository, producer)
	characteristicService := characteristic_service.NewService(characteristicRepository, productRepository, producer)
	productService := product_service.NewService(productRepository, s3Client, producer, categoryService, characteristicService)

	// Controllers
	productController := product.NewController(productService)
	categoryController := category.NewController(categoryService)
	characteristicController := characteristic.NewController(characteristicService)

	// Run server
	srv := server.NewServer(
		&server.Config{
			Host: cfg.Server.Host,
			Port: cfg.Server.Port,
			Controllers: []server.Controller{
				productController,
				categoryController,
				characteristicController,
			},
		},
		server.WithLogger,
		server.WithMetrics,
	)

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
