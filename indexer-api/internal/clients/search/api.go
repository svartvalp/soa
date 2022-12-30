package search

import (
	"context"
	"encoding/json"

	"github.com/soa/indexer-api/internal/config"
	"github.com/soa/indexer-api/internal/models"
	"github.com/svartvalp/soa/service/api"
)

const (
	updateProduct string = "update product"
	commitIndex   string = "commit index"
)

type (
	Client struct {
		api api.API
	}
)

func New(cfg *config.Config) *Client {
	return &Client{
		api: api.NewAPI(api.APICfg{
			Address: cfg.SearchAPI.Address,
			Handles: api.GetHandels(cfg.SearchAPI.Handles),
		}),
	}
}

func (s *Client) SendNewInfo(ctx context.Context, info []models.ProductInfo) error {
	body, err := json.Marshal(info)
	if err != nil {
		return err
	}

	_, err = s.api.DoRequest(ctx, updateProduct, body)
	if err != nil {
		return err
	}

	_, err = s.api.DoRequest(ctx, commitIndex, nil)
	if err != nil {
		return err
	}

	return nil
}
