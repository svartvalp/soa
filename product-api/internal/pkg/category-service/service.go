package category_service

import (
	"context"
	"encoding/json"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/kafka"
	"github.com/soa/product-api/internal/models"
)

type Service struct {
	repo        repository
	productRepo productRepository
	kafka       kafka.Producer
}

func NewService(repo repository, kafka kafka.Producer) *Service {
	return &Service{
		repo:  repo,
		kafka: kafka,
	}
}

func (s *Service) List(ctx context.Context) ([]*models.Category, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetCatsMap(ctx context.Context) (map[int64]models.Category, error) {
	res := make(map[int64]models.Category)
	cats, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, cat := range cats {
		res[cat.ID] = *cat
	}
	return res, err
}

func (s *Service) Create(ctx context.Context, in *dto.CreateCategoryReq) (int64, error) {
	id, err := s.repo.Create(ctx, &models.Category{
		Name:        in.Name,
		Description: in.Description,
		ParentID:    in.ParentID,
		Level:       in.Level,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) Update(ctx context.Context, cats *models.Category) error {
	err := s.repo.Update(ctx, cats)
	if err != nil {
		return err
	}

	products, err := s.productRepo.List(ctx, &models.ProductFilters{
		CategoryIDs: []int64{cats.ID},
	})
	if err != nil {
		return err
	}

	ids := make([]int64, 0, len(products))
	for _, product := range products {
		ids = append(ids, product.ID)
	}

	msg := kafka.Msg{
		Service: "productAPI",
		Type:    "UPDATE",
		IDs:     ids,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = s.kafka.Write(ctx, b)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	err = s.kafka.Write(ctx, []byte("productAPI"))
	if err != nil {
		return err
	}
	return nil
}
