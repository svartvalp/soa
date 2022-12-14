package catalog

import (
	"context"

	"github.com/soa/catalog-api/internal/models"
)

type (
	productClient interface {
		BrandList(context.Context) ([]string, error)
		CategoriesList(context.Context) ([]models.Category, error)
		GetProductsInfo(context.Context, []int64) ([]models.ProductInfo, error)
	}

	searchClient interface {
		GetProductIDs(context.Context, *models.Filter) ([]int64, error)
	}

	s3 interface {
		GetImage(ctx context.Context, filename string) (*models.Image, error)
	}
)
