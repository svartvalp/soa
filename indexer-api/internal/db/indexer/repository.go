package indexer

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/soa/indexer-api/internal/db"
	"github.com/soa/indexer-api/internal/models"
)

type Repository struct {
	db.Wrapper
}

func NewRepository(wp db.Wrapper) *Repository {
	return &Repository{
		wp,
	}
}

func (r Repository) List(ctx context.Context) ([]models.ProductInfo, error) {
	qb := db.PgQb().
		Select("*").
		From(fmt.Sprintf("%s AS p", models.ProductTableName)).
		OrderBy("id")

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

func (r Repository) Delete(ctx context.Context, ids []int64) error {
	qb := db.PgQb().
		Delete(models.ProductTableName)

	if len(ids) > 0 {
		qb = qb.Where(squirrel.Eq{"id": ids})
	}

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Create(ctx context.Context, infos []models.ProductInfo) error {
	qb := db.PgQb().Insert(models.ProductTableName).Columns(
		"id",
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
			info.ID,
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

func (r Repository) Update(ctx context.Context, in models.ProductInfo) error {
	qb := db.PgQb().
		Update(models.ProductTableName).
		SetMap(map[string]interface{}{
			"name":            in.Name,
			"category_id":     in.CategoryID,
			"price":           in.Price,
			"description":     in.Description,
			"brand":           in.Brand,
			"image":           in.Image,
			"characteristics": in.Characteristics,
			"categorys":       in.Categorys,
		}).Where(squirrel.Eq{"id": in.ID}).Suffix("RETURNING id")

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}
