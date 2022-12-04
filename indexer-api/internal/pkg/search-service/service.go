package search_service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/models"
)

type (
	Service struct {
		searchAPIAddress string
		searchPath       string
		requester        requester
	}
)

func NewService(cfg *config.Config, req requester) *Service {
	return &Service{
		searchAPIAddress: cfg.SearchAPI.Address,
		searchPath:       cfg.SearchAPI.Path,
		requester:        req,
	}
}

func (s *Service) SendNewInfo(ctx context.Context, info []models.ProductInfo) error {
	body, err := json.Marshal(info)
	_, err = s.requester.DoRequest(ctx, s.searchAPIAddress+s.searchPath, http.MethodPut, body)
	if err != nil {
		return err
	}
	return nil
}
