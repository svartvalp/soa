package product_service

import (
	"context"
)

type Service struct {
	repository repository
	productAPI productClient
	searchAPI  searchClient
}

func New(
	repository repository,
	productClient productClient,
	searchClient searchClient,
) *Service {
	return &Service{
		repository: repository,
		productAPI: productClient,
		searchAPI:  searchClient,
	}
}

func (s *Service) Regenerate(ctx context.Context) error {
	err := s.repository.Delete(ctx, nil)
	if err != nil {
		return err
	}

	info, err := s.productAPI.GetNewData(ctx, nil)
	if err != nil {
		return err
	}

	err = s.repository.Create(ctx, info)
	if err != nil {
		return err
	}

	return s.searchAPI.SendNewInfo(ctx, info)
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

	err = s.repository.Create(ctx, info)
	if err != nil {
		return err
	}

	repoInfo, err := s.repository.List(ctx)
	if err != nil {
		return err
	}

	return s.searchAPI.SendNewInfo(ctx, repoInfo)
}
