package catalog

import (
	"context"

	"github.com/soa/catalog-api/internal/models"
)

type (
	catalogService interface {
		BrandList(context.Context) ([]string, error)
		CategoryList(context.Context) ([]models.Category, error)
		ProductList(context.Context, *models.Filter) ([]models.ProductInfo, error)
	}
)
