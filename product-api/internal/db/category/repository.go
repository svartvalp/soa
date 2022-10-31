package category

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

func (r *Repository) List(ctx context.Context) ([]*models.Category, error) {
	qb := db.PgQb().Select("*").From(models.CategoryTableName)

	var res []*models.Category
	err := r.Selectx(ctx, &res, qb)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (r *Repository) Create(ctx context.Context, in *models.Category) (int64, error) {
	qb := db.PgQb().Insert(models.CategoryTableName).
		Columns(
			"name",
			"description",
			"parent_id",
			"level",
		).
		Values(
			in.Name,
			in.Description,
			in.ParentID,
			in.Level,
		).Suffix("RETURNING id")
	id, err := r.Execx(ctx, qb)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, in *models.Category) error {
	qb := db.PgQb().
		Update(models.CharacteristicTableName).
		SetMap(map[string]interface{}{
			"name":      in.Name,
			"ch_type":   in.Description,
			"parent_id": in.ParentID,
			"level":     in.Level,
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
		Delete(models.CategoryTableName).
		Where(squirrel.Eq{"id": id})

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}
