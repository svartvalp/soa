package search

import (
	"context"

	"github.com/soa/search-api/internal/models"
)

type (
	productService interface {
		List(context.Context, *models.Filter) ([]models.ProductInfo, error)
		Regenerate(context.Context, []models.ProductInfo) error
	}
)
