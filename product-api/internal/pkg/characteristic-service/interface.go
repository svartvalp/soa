package characteristic_service

import (
	"context"

	"github.com/soa/product-api/internal/models"
)

type repository interface {
	List(ctx context.Context) ([]*models.Characteristic, error)
	ProductCharacteristicList(ctx context.Context) ([]*models.ProductCharacteristic, error)
	Create(ctx context.Context, product *models.Characteristic) (int64, error)
	Update(ctx context.Context, in *models.Characteristic) error
	Delete(ctx context.Context, id int64) error
}
