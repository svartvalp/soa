package product_service

import (
	"context"
	"fmt"

	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/kafka"
	"github.com/soa/product-api/internal/models"
)

type Service struct {
	categoryService       categoryService
	characteristicService characteristicService
	repo                  repository
	kafka                 kafka.Producer
	s3                    s3
}

func NewService(repo repository, s3 s3, producer kafka.Producer, categoryService categoryService, characteristicService characteristicService) *Service {
	return &Service{
		categoryService:       categoryService,
		characteristicService: characteristicService,
		repo:                  repo,
		kafka:                 producer,
		s3:                    s3,
	}
}

func (s *Service) GetFullProductInfo(ctx context.Context) ([]models.FullProductInfo, error) {
	res := make([]models.FullProductInfo, 0, 100)
	products, err := s.List(ctx, nil)
	if err != nil {
		return nil, err
	}
	catsMap, err := s.categoryService.GetCatsMap(ctx)
	if err != nil {
		return nil, err
	}
	characteristicMap, err := s.characteristicService.GetCharacteristicMap(ctx)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		cats := getCategorys(nil, catsMap, product.CategoryID)
		characteristics := characteristicMap[product.ID]
		res = append(res, models.FullProductInfo{
			ID:              product.ID,
			CategoryID:      product.CategoryID,
			Price:           product.Price,
			Name:            product.Name,
			Description:     product.Description,
			Brand:           product.Brand,
			Image:           product.Image,
			Characteristics: characteristics,
			Categorys:       cats,
		})
	}
	return res, nil
}

func getCategorys(cats []models.Category, m map[int64]models.Category, id int64) []models.Category {
	if len(cats) == 0 {
		cats = make([]models.Category, 0, 100)
		if cat, ok := m[id]; ok {
			cats = append(cats, cat)
			return getCategorys(cats, m, 0)
		}
	} else if cat, ok := m[cats[len(cats)-1].ParentID]; ok {
		cats = append(cats, cat)
		return getCategorys(cats, m, 0)
	}
	return cats
}

func (s *Service) List(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error) {
	return s.repo.List(ctx, req)
}

func (s *Service) Create(ctx context.Context, product *dto.CreateProductReq) (int64, error) {
	id, err := s.repo.Create(ctx, &models.Product{
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Name,
		Brand:       product.Brand,
	})
	if err != nil {
		return 0, err
	}
	err = s.kafka.Write(ctx, []byte("productAPI"))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) UI(ctx context.Context, id int64, img *dto.Image) error {
	imageName := imageName(img.Name, id)
	err := s.s3.PutImage(ctx, img.Body, imageName)
	if err != nil {
		return err
	}
	err = s.repo.SetImage(ctx, imageName, id)
	if err != nil {
		return err
	}

	err = s.kafka.Write(ctx, []byte("productAPI"))
	if err != nil {
		return err
	}

	return nil
}

func imageName(name string, id int64) string {
	return fmt.Sprintf("%d-%s", id, name)
}

func (s *Service) Update(ctx context.Context, products *models.Product) error {
	err := s.repo.Update(ctx, products)
	if err != nil {
		return err
	}

	err = s.kafka.Write(ctx, []byte("productAPI"))
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
