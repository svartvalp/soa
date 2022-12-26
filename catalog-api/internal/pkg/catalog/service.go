package catalog

import (
	"context"

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

func (s *Service) ProductList(ctx context.Context, req *models.Filter) ([]models.ProductInfo, error) {
	return s.searchAPI.ProductList(ctx, req)
}

func New(productAPI productAPI, searchAPI searchAPI) *Service {
	return &Service{
		productAPI: productAPI,
		searchAPI:  searchAPI,
	}
}
