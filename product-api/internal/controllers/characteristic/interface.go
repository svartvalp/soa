package characteristic

import (
	"context"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type (
	characteristicService interface {
		List(ctx context.Context) ([]*models.Characteristic, error)
		Create(ctx context.Context, in *dto.CreateCharacteristicReq) (int64, error)
		Update(ctx context.Context, products *models.Characteristic) error
		Delete(ctx context.Context, id int64) error
	}
)
