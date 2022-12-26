package catalog

import (
	"context"

	"github.com/soa/catalog-api/internal/models"
)

type (
	productAPI interface {
		BrandList(context.Context) ([]string, error)
		CategoriesList(context.Context) ([]models.Category, error)
	}

	searchAPI interface {
		ProductList(context.Context, *models.Filter) ([]models.ProductInfo, error)
	}
)
