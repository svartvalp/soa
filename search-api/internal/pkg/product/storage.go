package product

import (
	"context"

	"github.com/soa/search-api/internal/models"
)

type Storage struct {
	s map[int64]models.ProductInfo
}

func (s *Storage) GetByIDs(ctx context.Context, ids []int64) ([]models.ProductInfo, error) {
	res := make([]models.ProductInfo, 0, len(ids))
	for _, id := range ids {
		if pr, ok := s.s[id]; ok {
			res = append(res, pr)
		}
	}
	return res, nil
}

func (s *Storage) Set(ctx context.Context, pr models.ProductInfo) error {
	s.s[pr.ID] = pr
	return nil
}

func New() *Storage {
	return &Storage{s: make(map[int64]models.ProductInfo)}
}
