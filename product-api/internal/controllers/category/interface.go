package category

import (
	"context"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type (
	categoryService interface {
		List(ctx context.Context) ([]*models.Category, error)
		Create(ctx context.Context, in *dto.CreateCategoryReq) (int64, error)
		Update(ctx context.Context, products *models.Category) error
		Delete(ctx context.Context, id int64) error
	}
)
