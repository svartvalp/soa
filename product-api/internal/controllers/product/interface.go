package product

import (
	"context"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type (
	productService interface {
		List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error)
		BrandList(ctx context.Context) ([]string, error)
		GetFullProductInfo(ctx context.Context, req *models.ProductFilters) ([]models.FullProductInfo, error)
		Create(ctx context.Context, product *dto.CreateProductReq) (int64, error)
		Update(ctx context.Context, products *models.Product) error
		Delete(ctx context.Context, id int64) error
		UI(ctx context.Context, id int64, img *dto.Image) error
	}
)
