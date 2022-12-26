package product

import (
	"context"
	"net/http"

	"github.com/soa/catalog-api/internal/config"
	"github.com/soa/catalog-api/internal/models"
	"github.com/soa/catalog-api/internal/requester"
)

const (
	brandListURL    = "product/brand/list"
	categoryListURL = "category/list"
)

type ProductAPI struct {
	r requester.Requester

	address string
}

func (p *ProductAPI) BrandList(ctx context.Context) ([]string, error) {
	resp, err := p.r.DoRequest(ctx, p.address+brandListURL, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]string](resp)
}

func (p *ProductAPI) CategoriesList(ctx context.Context) ([]models.Category, error) {
	resp, err := p.r.DoRequest(ctx, p.address+categoryListURL, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return requester.UnmarshalBody[[]models.Category](resp)
}

func New(r requester.Requester, cfg *config.Config) *ProductAPI {
	return &ProductAPI{
		r:       r,
		address: cfg.ProductAPI.Address,
	}
}
