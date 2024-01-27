package rank

import (
	"estore/internal/types"
	"time"
)

func TableName() string {
	return "t_rank"
}

type Rank struct {
	UUID       string           `gorm:"type:varchar(40);primaryKey;not null"`
	Club       string           `gorm:"type:varchar(40);not null"`
	Name       string           `gorm:"type:varchar(40);not null"`
	Conditions types.Conditions `gorm:"type:json;not null"`
	Priority   int32            `gorm:"type:tinyint;not null"`
	CreatedAt  time.Time        `gorm:"->;type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (Rank) TableName() string {
	return TableName()
}

func (r Rank) GetUUID() string {
	return r.UUID
}
