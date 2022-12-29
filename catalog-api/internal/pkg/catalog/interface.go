package catalog

import (
	"context"

	"github.com/soa/catalog-api/internal/models"
)

type (
	productAPI interface {
		BrandList(context.Context) ([]string, error)
		CategoriesList(context.Context) ([]models.Category, error)
		GetProductsInfo(context.Context, []int64) ([]models.ProductInfo, error)
	}

	searchAPI interface {
		GetProductIDs(context.Context, *models.Filter) ([]int64, error)
	}

	s3 interface {
		GetImage(ctx context.Context, filename string) (*models.Image, error)
	}
)
