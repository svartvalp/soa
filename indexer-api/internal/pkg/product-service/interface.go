package product_service

import (
	"context"

	"github.com/soa/indexer-api/internal/models"
)

type (
	repository interface {
		List(ctx context.Context) ([]models.ProductInfo, error)
		Delete(context.Context, []int64) error
		Update(context.Context, models.ProductInfo) error
		Create(context.Context, []models.ProductInfo) error
	}

	productClient interface {
		GetNewData(context.Context, []int64) ([]models.ProductInfo, error)
	}

	searchClient interface {
		SendNewInfo(context.Context, []models.ProductInfo) error
	}
)
