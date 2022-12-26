package product_service

import (
	"context"

	"github.com/soa/product-api/internal/models"
)

type (
	repository interface {
		List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error)
		BrandList(ctx context.Context) ([]string, error)
		Create(ctx context.Context, product *models.Product) (int64, error)
		Update(ctx context.Context, in *models.Product) error
		Delete(ctx context.Context, id int64) error
		SetImage(ctx context.Context, img string, id int64) error
	}
	s3 interface {
		PutImage(ctx context.Context, img []byte, key string) error
		GetImage(ctx context.Context, key string) ([]byte, error)
	}

	categoryService interface {
		GetCatsMap(ctx context.Context) (map[int64]models.Category, error)
	}

	characteristicService interface {
		GetCharacteristicMap(ctx context.Context) (map[int64][]models.ProductCharacteristic, error)
	}
)
