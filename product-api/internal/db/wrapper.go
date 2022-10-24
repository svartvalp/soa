package db

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/soa/product-api/internal/config"
)

// PgQb sets placeholder format
func PgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}

func NewPGConnection(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, cfg.DatabaseDsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
