package dao

import (
	"gorm.io/gorm"
)

type PrizeModel struct {
	// ID is the unique identifier (Auto-incremented: 1, 2, 3...)
	ID int64 `gorm:"primaryKey;autoIncrement"`

	// Name of the prize (e.g., "iPhone 16 Pro")
	Name string `gorm:"type:varchar(255);not null;index"`

	// Pic is the prize image URL for frontend display
	Pic string `gorm:"type:varchar(512)"`

	// Link is a URL for prize details or redemption
	Link string `gorm:"type:varchar(512)"`

	// Type category: 1=Physical (requires shipping), 2=Virtual (coupon/code), 3=Credit/Cash
	Type int32 `gorm:"type:tinyint;default:0"`

	// Data stores JSON configuration (e.g., {"coupon_code": "XYZ123"})
	Data string `gorm:"type:text"`

	// Total is the initial quantity allocated for the event
	Total int64 `gorm:"not null;default:0"`

	// Left is the current stock. Decremented by 1 per draw.
	Left int64 `gorm:"not null;default:0"`

	// IsUse indicates availability: 1=Enabled (active), 0=Disabled (hidden/paused)
	IsUse int32 `gorm:"type:tinyint;default:1;index"`

	// Probability value (weight). Sum of all prizes usually equals 10000.
	Probability int64 `gorm:"default:0"`

	// ProbabilityMax/Min define the lottery winning range.
	// A user wins this prize if: RandomNumber >= Min AND RandomNumber < Max
	ProbabilityMax int64 `gorm:"default:0;index"`
	ProbabilityMin int64 `gorm:"default:0;index"`

	// Timestamps managed automatically by GORM in milliseconds
	Utime int64 `gorm:"autoUpdateTime:milli"`
	Ctime int64 `gorm:"autoCreateTime:milli"`
}

// TableName specifies the table name for PrizeModel
func (PrizeModel) TableName() string {
	return "prizes"
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
