package product_api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/models"
	"github.com/soa/indexer-api/internal/requester"
)

const (
	productInfoListURL = "product/full-info"
)

type (
	ProductAPI struct {
		productAPIAddress string
		productPath       string
		requester         *requester.Requester
	}
)

func New(cfg *config.Config, req *requester.Requester) *ProductAPI {
	return &ProductAPI{
		productAPIAddress: cfg.ProductAPI.Address,
		requester:         req,
	}
}

func (s *ProductAPI) GetNewData(ctx context.Context, ids []int64) ([]models.ProductInfo, error) {
	b, err := json.Marshal(models.ProductFilters{
		IDs: ids,
	})
	if err != nil {
		return nil, err
	}

	resp, err := s.requester.DoRequest(ctx, s.productAPIAddress+productInfoListURL, http.MethodGet, b)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]models.ProductInfo](resp)
}
