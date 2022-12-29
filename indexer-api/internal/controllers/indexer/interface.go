package indexer

import (
	"context"
)

type indexerService interface {
	Regenerate(ctx context.Context) error
}
