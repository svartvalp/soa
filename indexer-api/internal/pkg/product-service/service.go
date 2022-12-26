package product_service

import (
	"context"
)

type Service struct {
	repository repository
	productAPI productAPI
	searchAPI  searchAPI
}

func New(
	repository repository,
	productAPI productAPI,
	searchAPI searchAPI,
) *Service {
	return &Service{
		repository: repository,
		productAPI: productAPI,
		searchAPI:  searchAPI,
	}
}

func (s *Service) ProductAPIDeleteIvent(ctx context.Context, ids []int64) error {
	err := s.repository.Delete(ctx, ids)
	if err != nil {
		return err
	}

	repoInfo, err := s.repository.List(ctx)
	if err != nil {
		return err
	}

	return s.searchAPI.SendNewInfo(ctx, repoInfo)
}

func (s *Service) ProductAPIUpdateIvent(ctx context.Context, ids []int64) error {
	info, err := s.productAPI.GetNewData(ctx, ids)
	if err != nil {
		return err
	}

	err = s.repository.Update(ctx, info[0])
	if err != nil {
		return err
	}

	repoInfo, err := s.repository.List(ctx)
	if err != nil {
		return err
	}

	return s.searchAPI.SendNewInfo(ctx, repoInfo)
}

func (s *Service) ProductAPICreateIvent(ctx context.Context, ids []int64) error {
	info, err := s.productAPI.GetNewData(ctx, ids)
	if err != nil {
		return err
	}

	if len(info) != 1 {
		return nil
	}

	err = s.repository.Create(ctx, info[0])
	if err != nil {
		return err
	}

	repoInfo, err := s.repository.List(ctx)
	if err != nil {
		return err
	}

	return s.searchAPI.SendNewInfo(ctx, repoInfo)
}
