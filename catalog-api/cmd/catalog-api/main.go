package main

import (
	"log"

	"github.com/soa/catalog-api/internal/config"
	catalog_controller "github.com/soa/catalog-api/internal/controllers/catalog"
	"github.com/soa/catalog-api/internal/pkg/catalog"
	"github.com/soa/catalog-api/internal/requester"
	product_api "github.com/soa/catalog-api/internal/requester/product"
	search_api "github.com/soa/catalog-api/internal/requester/search"
	"github.com/soa/catalog-api/internal/server"
)

func main() {
	cfg, err := config.NewConfig("internal/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	requester := requester.NewRequester()

	// APIs
	productAPI := product_api.New(requester, cfg)
	searchAPI := search_api.New(requester, cfg)

	// Services
	catalogService := catalog.New(productAPI, searchAPI)

	// Controllers
	catalogController := catalog_controller.New(catalogService)

	srv := server.NewServer(cfg, catalogController)

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
