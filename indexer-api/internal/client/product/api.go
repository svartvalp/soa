package product

import (
	"context"
	"encoding/json"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/models"
	"github.com/svartvalp/soa/service/api"
)

const (
	fullInfo = "full info"
)

type (
	ProductAPI struct {
		api api.API
	}
)

func New(cfg *config.Config) *ProductAPI {
	return &ProductAPI{
		api: api.NewAPI(api.APICfg{
			Address: cfg.ProductAPI.Address,
			Handles: api.GetHandels(cfg.ProductAPI.Handles),
		}),
	}
}

func (s *ProductAPI) GetNewData(ctx context.Context, ids []int64) ([]models.ProductInfo, error) {
	b, err := json.Marshal(models.ProductFilters{
		IDs: ids,
	})
	if err != nil {
		return nil, err
	}

	resp, err := s.api.DoRequest(ctx, fullInfo, b)
	if err != nil {
		return nil, err
	}

	return api.UnmarshalBody[[]models.ProductInfo](resp)
}
