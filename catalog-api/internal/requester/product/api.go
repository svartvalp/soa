package product

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/soa/catalog-api/internal/config"
	"github.com/soa/catalog-api/internal/models"
	"github.com/soa/catalog-api/internal/requester"
)

const (
	brandListURL    = "product/brand/list"
	productInfoURL  = "product/full-info"
	categoryListURL = "category/list"
)

type ProductAPI struct {
	r requester.Requester

	address string
}

func New(r requester.Requester, cfg *config.Config) *ProductAPI {
	return &ProductAPI{
		r:       r,
		address: cfg.ProductAPI.Address,
	}
}

func (p *ProductAPI) BrandList(ctx context.Context) ([]string, error) {
	resp, err := p.r.DoRequest(ctx, p.address+brandListURL, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]string](resp)
}

func (p *ProductAPI) GetProductsInfo(ctx context.Context, ids []int64) ([]models.ProductInfo, error) {
	b, err := json.Marshal(&models.FullProductFilters{
		IDs: ids,
	})
	if err != nil {
		return nil, err
	}

	resp, err := p.r.DoRequest(ctx, p.address+productInfoURL, http.MethodPost, b)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]models.ProductInfo](resp)
}

func (p *ProductAPI) CategoriesList(ctx context.Context) ([]models.Category, error) {
	resp, err := p.r.DoRequest(ctx, p.address+categoryListURL, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]models.Category](resp)
}
