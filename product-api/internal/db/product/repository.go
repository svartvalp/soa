package product

import (
	"context"
	"fmt"

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

func (r *Repository) List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error) {
	qb := applyFilter(req)
	var res []*models.Product
	err := r.Selectx(ctx, &res, qb)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func applyFilter(req *models.ProductFilters) squirrel.SelectBuilder {
	qb := db.PgQb().
		Select(
			"p.id",
			"p.name",
			"p.description",
			"p.category_id",
			"p.brand",
			"p.price",
			"p.image",
		).
		From(fmt.Sprintf("%s AS p", models.ProductTableName)).
		LeftJoin(fmt.Sprintf("%v pc on p.id = pc.product_id", models.ProductCharacteristicTableName))
	if req != nil {
		if len(req.CategoryIDs) > 0 {
			qb = qb.Where(squirrel.Eq{"p.category_id": req.CategoryIDs})
		}
		if len(req.Names) > 0 {
			qb = qb.Where(squirrel.Eq{"p.name": req.Names})
		}
		if len(req.ProductCharacteristicIDs) > 0 {
			qb = qb.Where(squirrel.Eq{"pc.id": req.ProductCharacteristicIDs})
		}
	}
	qb = qb.OrderBy("id")
	return qb
}

func (r *Repository) Create(ctx context.Context, product *models.Product) (int64, error) {
	qb := db.PgQb().Insert(models.ProductTableName).
		Columns(
			"category_id",
			"price",
			"name",
			"description",
			"brand",
			"image",
		).
		Values(
			product.CategoryID,
			product.Price,
			product.Name,
			product.Description,
			product.Brand,
			product.Image,
		).Suffix("RETURNING id")
	id, err := r.Execx(ctx, qb)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, in *models.Product) error {
	qb := db.PgQb().
		Update(models.ProductTableName).
		SetMap(map[string]interface{}{
			"category_id": in.CategoryID,
			"price":       in.Price,
			"name":        in.Name,
			"description": in.Description,
			"brand":       in.Brand,
			"image":       in.Image,
		}).Where(squirrel.Eq{"id": in.ID}).Suffix("RETURNING id")

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}

func (r *Repository) SetImage(ctx context.Context, img string, id int64) error {
	qb := db.PgQb().
		Update(models.ProductTableName).
		Set(
			"image", img,
		).Where(squirrel.Eq{"id": id}).Suffix("RETURNING id")

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	qb := db.PgQb().
		Delete(models.ProductTableName).
		Where(squirrel.Eq{"id": id})

	_, err := r.Execx(ctx, qb)
	if err != nil {
		return err
	}
	return err
}
