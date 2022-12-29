package db

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/soa/product-api/internal/config"
)

type (
	Wrapper interface {
		Selectx(ctx context.Context, dest interface{}, qb squirrel.SelectBuilder) error
		Execx(ctx context.Context, qb squirrel.Sqlizer) (int64, error)
	}
	wrapper struct {
		pg *pgxpool.Pool
	}
)

// PgQb sets placeholder format
func PgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}

func (w *wrapper) Selectx(ctx context.Context, dest interface{}, qb squirrel.SelectBuilder) error {
	stmt, args, err := qb.ToSql()
	if err != nil {
		return err
	}
	rows, err := w.pg.Query(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (w *wrapper) Execx(ctx context.Context, qb squirrel.Sqlizer) (int64, error) {
	stmt, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	err = w.pg.QueryRow(ctx, stmt, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func NewWrapper(ctx context.Context, cfg *config.Config) (Wrapper, error) {
	pool, err := pgxpool.Connect(ctx, cfg.DatabaseDsn)
	if err != nil {
		return nil, err
	}
	return &wrapper{pool}, nil
}
