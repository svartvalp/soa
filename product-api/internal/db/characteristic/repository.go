package characteristic

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/soa/product-api/internal/db"
	"github.com/soa/product-api/internal/models"
)

type Repository struct {
	db.Wrapper
}

func NewRepository(wp db.Wrapper) *Repository {
	return &Repository{
		wp,
	}
}

func (r *Repository) List(ctx context.Context) ([]*models.Characteristic, error) {
	qb := db.PgQb().Select("*").From(models.CharacteristicTableName)

	var res []*models.Characteristic
	err := r.Selectx(ctx, &res, qb)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (r *Repository) Create(ctx context.Context, in *models.Characteristic) (int64, error) {
	qb := db.PgQb().Insert(models.CharacteristicTableName).
		Columns(
			"name",
			"ch_type",
			"description",
		).
		Values(
			in.Name,
			in.ChType,
			in.Descriptions,
		).Suffix("RETURNING id")
	id, err := r.Execx(ctx, qb)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, in *models.Characteristic) error {
	qb := db.PgQb().
		Update(models.CharacteristicTableName).
		SetMap(map[string]interface{}{
			"name":        in.Name,
			"ch_type":     in.ChType,
			"description": in.Descriptions,
		}).Where(squirrel.Eq{"id": in.ID}).
		Suffix("RETURNING id")

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	qb := db.PgQb().
		Delete(models.CharacteristicTableName).
		Where(squirrel.Eq{"id": id})

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}
