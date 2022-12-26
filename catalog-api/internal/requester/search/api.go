package search

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/soa/catalog-api/internal/config"
	"github.com/soa/catalog-api/internal/models"
	"github.com/soa/catalog-api/internal/requester"
)

const (
	productListURL = "product/list"
)

type SearchAPI struct {
	r requester.Requester

	address string
}

func (p *SearchAPI) GetProductIDs(ctx context.Context, filter *models.Filter) ([]int64, error) {
	b, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	resp, err := p.r.DoRequest(ctx, p.address+productListURL, http.MethodPost, b)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]int64](resp)
}

func New(r requester.Requester, cfg *config.Config) *SearchAPI {
	return &SearchAPI{
		r:       r,
		address: cfg.SearchAPI.Address,
	}
}
