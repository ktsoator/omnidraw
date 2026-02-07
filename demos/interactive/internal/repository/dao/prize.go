package dao

import (
	"gorm.io/gorm"
)

type PrizeModel struct {
	Id    int64  `gorm:"primaryKey,autoIncrement"`
	Name  string `gorm:"type:varchar(255)"`
	Count int    `gorm:"type:int"`
	Ext   string `gorm:"type:text"`
	Utime int64
	Ctime int64
}

type PrizeDAO struct {
	db *gorm.DB
}

func NewPrizeDAO(db *gorm.DB) *PrizeDAO {
	return &PrizeDAO{db: db}
}

func (dao *PrizeDAO) Insert(p PrizeModel) error {
	return dao.db.Create(&p).Error
}
