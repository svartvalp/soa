package search_service

import (
	"context"

	"github.com/soa/search-api/internal/models"
)

type (
	repository interface {
		List(context.Context, *models.Filter) ([]models.ProductInfo, error)
		DeleteAll(ctx context.Context) error
		Create(context.Context, []models.ProductInfo) error
	}
	prStorage interface {
		GetByIDs(ctx context.Context, ids []int64) ([]models.ProductInfo, error)
		Set(ctx context.Context, pr models.ProductInfo) error
	}
)
