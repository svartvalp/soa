package characteristic_service

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

func NewService(repo repository, productRepo productRepository, kafka kafka.Producer) *Service {
	return &Service{
		repo:        repo,
		productRepo: productRepo,
		kafka:       kafka,
	}
}

func (s *Service) List(ctx context.Context) ([]*models.Characteristic, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetCharacteristicMap(ctx context.Context) (map[int64][]models.ProductCharacteristic, error) {
	res := make(map[int64][]models.ProductCharacteristic)
	chars, err := s.repo.ProductCharacteristicList(ctx)
	if err != nil {
		return nil, err
	}

	for _, char := range chars {
		res[char.ProductId] = append(res[char.ProductId], *char)
	}

	return res, err
}

func (s *Service) Create(ctx context.Context, in *dto.CreateCharacteristicReq) (int64, error) {
	id, err := s.repo.Create(ctx, &models.Characteristic{
		Name:        in.Name,
		ChType:      in.ChType,
		Description: in.Description,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) Update(ctx context.Context, char *models.Characteristic) error {
	err := s.repo.Update(ctx, char)
	if err != nil {
		return err
	}

	products, err := s.productRepo.List(ctx, &models.ProductFilters{
		CharacteristicIDs: []int64{char.ID},
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
