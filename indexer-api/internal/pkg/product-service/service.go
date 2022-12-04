package product_service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/models"
)

type (
	Service struct {
		productAPIAddress string
		productPath       string
		requester         requester
	}
)

func NewService(cfg *config.Config, req requester) *Service {
	return &Service{
		productAPIAddress: cfg.ProductAPI.Address,
		productPath:       cfg.ProductAPI.Path,
		requester:         req,
	}
}

func (s *Service) GetNewData(ctx context.Context) ([]models.ProductInfo, error) {
	resp, err := s.requester.DoRequest(ctx, s.productAPIAddress+s.productPath, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var info []models.ProductInfo
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
