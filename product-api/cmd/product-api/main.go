package main

import (
	"context"

	"github.com/soa/product-api/internal/config"
	"github.com/soa/product-api/internal/controllers/category"
	"github.com/soa/product-api/internal/controllers/characteristic"
	"github.com/soa/product-api/internal/controllers/product"
	"github.com/soa/product-api/internal/db"
	category_repo "github.com/soa/product-api/internal/db/category-repo"
	characteristic_repo "github.com/soa/product-api/internal/db/characteristic-repo"
	product_repo "github.com/soa/product-api/internal/db/product-repo"
	category_service "github.com/soa/product-api/internal/pkg/category-service"
	characteristic_service "github.com/soa/product-api/internal/pkg/characteristic-service"
	product_service "github.com/soa/product-api/internal/pkg/product-service"
	"github.com/soa/product-api/internal/server"
)

func main() {
	ctx := context.Background()

	// Config
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		return
	}

	conn, err := db.NewPGConnection(ctx, cfg)

	// Repository
	productRepository := product_repo.NewRepository(conn)
	categoryRepository := category_repo.NewRepository(conn)
	characteristicRepository := characteristic_repo.NewRepository(conn)

	// Services
	productService := product_service.NewService(productRepository)
	categoryService := category_service.NewService(categoryRepository)
	characteristicService := characteristic_service.NewService(characteristicRepository)

	// Controllers
	productController := product.NewController(productService)
	categoryController := category.NewController(categoryService)
	characteristicController := characteristic.NewController(characteristicService)

	srv := server.NewServer(
		cfg,
		productController,
		categoryController,
		characteristicController,
	)

	if err = srv.Run(); err != nil {
		return
	}
}
