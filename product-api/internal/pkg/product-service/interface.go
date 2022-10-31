package product_service

import (
	"context"

	"github.com/soa/product-api/internal/models"
)

type (
	repository interface {
		List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error)
		Create(ctx context.Context, product *models.Product) (int64, error)
		Update(ctx context.Context, in *models.Product) error
		Delete(ctx context.Context, id int64) error
	}
	s3 interface {
		PutImage(ctx context.Context, img []byte, key string) error
		GetImage(ctx context.Context, key string) ([]byte, error)
	}
)
