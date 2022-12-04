package search_service

import (
	"context"

	"github.com/soa/search-api/internal/models"
)

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context, req *models.Filter) ([]models.ProductInfo, error) {
	res, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Regenerate(ctx context.Context, products []models.ProductInfo) error {
	err := s.repo.DeleteAll(ctx)
	if err != nil {
		return err
	}
	err = s.repo.Create(ctx, products)
	if err != nil {
		return err
	}
	return nil
}
