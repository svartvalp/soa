package search_api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/models"
)

const (
	productUpdateURL = "product/update"
)

type (
	SearchAPI struct {
		searchAPIAddress string
		requester        requester
	}
)

func New(cfg *config.Config, req requester) *SearchAPI {
	return &SearchAPI{
		searchAPIAddress: cfg.SearchAPI.Address,
		requester:        req,
	}
}

func (s *SearchAPI) SendNewInfo(ctx context.Context, info []models.ProductInfo) error {
	body, err := json.Marshal(info)
	_, err = s.requester.DoRequest(ctx, s.searchAPIAddress+productUpdateURL, http.MethodPut, body)
	if err != nil {
		return err
	}
	return nil
}
