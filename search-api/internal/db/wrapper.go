package db

import (
	"context"

	"github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v4/pgxpool"
	"github.com/randallmlough/pgxscan"
	"github.com/soa/search-api/internal/config"
)

type (
	Wrapper interface {
		Selectx(ctx context.Context, dest interface{}, qb squirrel.SelectBuilder) error
		Execx(ctx context.Context, qb squirrel.Sqlizer) (int64, error)
	}
	wrapper struct {
		pg *pgx.Pool
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

	return pgxscan.NewScanner(rows).Scan(dest)
}

func (w *wrapper) Execx(ctx context.Context, qb squirrel.Sqlizer) (int64, error) {
	stmt, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}
	rows, err := w.pg.Query(ctx, stmt, args...)
	if err != nil {
		return 0, err
	}

	return rows.CommandTag().RowsAffected(), nil
}

func NewWrapper(ctx context.Context, cfg *config.Config) (Wrapper, error) {
	conn, err := pgx.Connect(ctx, cfg.DatabaseDsn)
	if err != nil {
		return nil, err
	}
	return &wrapper{conn}, nil
}
