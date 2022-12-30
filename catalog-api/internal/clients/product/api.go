package product

import (
	"context"
	"encoding/json"

	"github.com/soa/catalog-api/internal/config"
	"github.com/soa/catalog-api/internal/models"
	"github.com/svartvalp/soa/service/api"
)

const (
	brandList    = "product list"
	productInfo  = "full info"
	categoryList = "category list"
)

type Client struct {
	api api.API
}

func New(cfg *config.Config) *Client {
	return &Client{
		api: api.NewAPI(api.APICfg{
			Address: cfg.ProductAPI.Address,
			Handles: api.GetHandels(cfg.ProductAPI.Handles),
		}),
	}
}

func (c *Client) BrandList(ctx context.Context) ([]string, error) {
	resp, err := c.api.DoRequest(ctx, brandList, nil)
	if err != nil {
		return nil, err
	}

	return api.UnmarshalBody[[]string](resp)
}

func (c *Client) GetProductsInfo(ctx context.Context, ids []int64) ([]models.ProductInfo, error) {
	b, err := json.Marshal(&models.FullProductFilters{
		IDs: ids,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.api.DoRequest(ctx, productInfo, b)
	if err != nil {
		return nil, err
	}

	return api.UnmarshalBody[[]models.ProductInfo](resp)
}

func (c *Client) CategoriesList(ctx context.Context) ([]models.Category, error) {
	resp, err := c.api.DoRequest(ctx, categoryList, nil)
	if err != nil {
		return nil, err
	}

	return api.UnmarshalBody[[]models.Category](resp)
}
