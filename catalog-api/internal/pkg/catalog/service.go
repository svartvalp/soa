package catalog

import (
	"context"
	"log"

	"github.com/soa/catalog-api/internal/models"
)

type Service struct {
	productAPI productAPI
	searchAPI  searchAPI
}

func (s *Service) BrandList(ctx context.Context) ([]string, error) {
	return s.productAPI.BrandList(ctx)
}

func (s *Service) CategoryList(ctx context.Context) ([]models.Category, error) {
	return s.productAPI.CategoriesList(ctx)
}

func (s *Service) GetProducts(ctx context.Context, req *models.Filter) ([]models.ProductInfo, error) {
	ids, err := s.searchAPI.GetProductIDs(ctx, req)
	if err != nil {
		log.Printf("GetProducts: %v", err)
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	return s.productAPI.GetProductsInfo(ctx, ids)
}

func New(productAPI productAPI, searchAPI searchAPI) *Service {
	return &Service{
		productAPI: productAPI,
		searchAPI:  searchAPI,
	}
}
