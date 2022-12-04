package kafka

import (
	"context"

	"github.com/soa/indexer-api/internal/models"
)

type (
	Consumer interface {
		Start(context.Context)
	}

	productService interface {
		GetNewData(context.Context) ([]models.ProductInfo, error)
	}

	searchService interface {
		SendNewInfo(context.Context, []models.ProductInfo) error
	}

	handle func(context.Context) error

	handleMap map[string]handle
)
