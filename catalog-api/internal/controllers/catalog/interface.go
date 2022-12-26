package catalog

import (
	"context"

	"github.com/soa/catalog-api/internal/models"
)

type (
	catalogService interface {
		BrandList(context.Context) ([]string, error)
		CategoryList(context.Context) ([]models.Category, error)
		GetProducts(context.Context, *models.Filter) ([]models.ProductInfo, error)
	}
)
