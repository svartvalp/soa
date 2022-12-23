package search

import (
	"context"

	"github.com/soa/search-api/internal/models"
)

type (
	productService interface {
		List(context.Context, *models.Filter) ([]int64, error)
		Regenerate(context.Context, []models.ProductInfo) error
		SwapIndex(ctx context.Context) error
	}
)
