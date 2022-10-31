package category_service

import (
	"context"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]*models.Category, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, in *dto.CreateCategoryReq) (int64, error) {
	return s.repo.Create(ctx, &models.Category{
		Name:        in.Name,
		Description: in.Description,
		ParentID:    in.ParentID,
		Level:       in.Level,
	})
}

func (s *Service) Update(ctx context.Context, products *models.Category) error {
	return s.repo.Update(ctx, products)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
