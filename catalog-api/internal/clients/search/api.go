package search

import (
	"context"
	"encoding/json"

	"github.com/soa/catalog-api/internal/config"
	"github.com/soa/catalog-api/internal/models"
	"github.com/svartvalp/soa/service/api"
)

const (
	productIDs = "product ids"
)

type Client struct {
	api api.API
}

func New(cfg *config.Config) *Client {
	return &Client{
		api: api.NewAPI(api.APICfg{
			Address: cfg.SearchAPI.Address,
			Handles: api.GetHandels(cfg.SearchAPI.Handles),
		}),
	}
}

func (c *Client) GetProductIDs(ctx context.Context, filter *models.Filter) ([]int64, error) {
	b, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	resp, err := c.api.DoRequest(ctx, productIDs, b)
	if err != nil {
		return nil, err
	}

	return api.UnmarshalBody[[]int64](resp)
}
