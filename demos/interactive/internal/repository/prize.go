package repository

import (
	"context"
	"interactive/internal/domain"
	"interactive/internal/repository/dao"
	"time"
)

type PrizeRepository struct {
	dao *dao.PrizeDAO
}

func NewPrizeRepository(dao *dao.PrizeDAO) *PrizeRepository {
	return &PrizeRepository{dao: dao}
}

func (repo *PrizeRepository) Create(ctx context.Context, p domain.Prize) error {
	now := time.Now().UnixMilli()
	return repo.dao.Insert(dao.PrizeModel{
		Name:  p.Name,
		Count: p.Count,
		Ctime: now,
		Utime: now,
	})
}
