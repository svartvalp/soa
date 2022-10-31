package product_service

import (
	"context"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type Service struct {
	repo repository
	s3   s3
}

func NewService(repo repository, s3 s3) *Service {
	return &Service{
		repo: repo,
		s3:   s3,
	}
}

func (s *Service) List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error) {
	return s.repo.List(ctx, req)
}

func (s *Service) Create(ctx context.Context, product *dto.CreateProductReq) (int64, error) {
	return s.repo.Create(ctx, &models.Product{
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Name,
		Brand:       product.Brand,
		// Image:       imagePath,
	})
}

func (s *Service) UI(ctx context.Context, img *dto.Image) error {
	return s.s3.PutImage(ctx, img.Body, img.Name)
}

func (s *Service) Update(ctx context.Context, products *models.Product) error {
	return s.repo.Update(ctx, products)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
