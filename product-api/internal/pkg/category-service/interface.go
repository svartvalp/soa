package category_service

import (
	"context"

	"github.com/soa/product-api/internal/models"
)

type (
	repository interface {
		List(ctx context.Context) ([]*models.Category, error)
		Create(ctx context.Context, product *models.Category) (int64, error)
		Update(ctx context.Context, in *models.Category) error
		Delete(ctx context.Context, id int64) error
	}

	productRepository interface {
		List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error)
	}
)
