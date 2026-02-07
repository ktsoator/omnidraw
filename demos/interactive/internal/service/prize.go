package service

import (
	"context"
	"interactive/internal/domain"
	"interactive/internal/repository"
)

type PrizeService struct {
	repo *repository.PrizeRepository
}

func NewPrizeService(repo *repository.PrizeRepository) *PrizeService {
	return &PrizeService{repo: repo}
}

func (s *PrizeService) AddPrize(ctx context.Context, p []domain.Prize) error {
	for _, v := range p {
		if err := s.repo.Create(ctx, v); err != nil {
			return err
		}
	}
	return nil
}
