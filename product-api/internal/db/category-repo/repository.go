package category_repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/soa/product-api/internal/db"
	"github.com/soa/product-api/internal/models"
)

type Repository struct {
	*pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		conn,
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]models.Product, error) {
	qb := db.PgQb().Select("*").From(models.ProductTableName)
	stmt, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Query(ctx, stmt, args)
	if err != nil {
		return nil, err
	}

	var res []models.Product

	for rows.Next() {
		var r models.Product
		err := rows.Scan()
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, r)
	}

	return res, nil
}
