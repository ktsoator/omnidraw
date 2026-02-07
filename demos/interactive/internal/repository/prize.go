package repository

import (
	"context"
	"interactive/internal/domain"
	"interactive/internal/repository/dao"
)

type PrizeRepository struct {
	dao *dao.PrizeDAO
}

func NewPrizeRepository(dao *dao.PrizeDAO) *PrizeRepository {
	return &PrizeRepository{dao: dao}
}

func (repo *PrizeRepository) Create(ctx context.Context, p domain.Prize) error {
	return repo.dao.Insert(dao.PrizeModel{
		ID:             p.ID,
		Name:           p.Name,
		Pic:            p.Pic,
		Link:           p.Link,
		Type:           p.Type,
		Data:           p.Data,
		Total:          p.Total,
		Left:           p.Left,
		IsUse:          p.IsUse,
		Probability:    p.Probability,
		ProbabilityMax: p.ProbabilityMax,
		ProbabilityMin: p.ProbabilityMin,
	})
}
