package search

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/soa/search-api/internal/db"
	"github.com/soa/search-api/internal/models"
)

type Repository struct {
	db.Wrapper
}

func NewRepository(wp db.Wrapper) *Repository {
	return &Repository{
		wp,
	}
}

func (r Repository) List(ctx context.Context, filter *models.Filter) ([]models.ProductInfo, error) {
	qb := applyFilter(filter)
	var res []models.ProductInfo
	err := r.Selectx(ctx, &res, qb)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func applyFilter(req *models.Filter) squirrel.SelectBuilder {
	qb := db.PgQb().
		Select("*").
		From(fmt.Sprintf("%s AS p", models.ProductTableName))

	if req != nil {
		// if len(req.CategoryIDs) > 0 {
		// 	qb = qb.Where(squirrel.Eq{"p.category_id": req.CategoryIDs})
		// }
		// if len(req.Names) > 0 {
		// 	qb = qb.Where(squirrel.Eq{"p.name": req.Names})
		// }
		// if len(req.ProductCharacteristicIDs) > 0 {
		// 	qb = qb.Where(squirrel.Eq{"pc.id": req.ProductCharacteristicIDs})
		// }
	}
	qb = qb.OrderBy("id")
	return qb
}

func (r Repository) DeleteAll(ctx context.Context) error {
	qb := db.PgQb().Delete(models.ProductTableName)
	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Create(ctx context.Context, infos []models.ProductInfo) error {
	if len(infos) == 0 {
		return nil
	}

	qb := db.PgQb().Insert(models.ProductTableName).Columns(
		"name",
		"category_id",
		"price",
		"description",
		"brand",
		"image",
		"characteristics",
		"categorys",
	)

	for _, info := range infos {
		qb = qb.Values(
			info.Name,
			info.CategoryID,
			info.Price,
			info.Description,
			info.Brand,
			info.Image,
			info.Characteristics,
			info.Categorys,
		)
	}
	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return nil
}
