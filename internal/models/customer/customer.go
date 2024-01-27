package customer

import (
	"estore/internal/types"
	"time"
)

func TableName() string {
	return "t_customer"
}

type Customer struct {
	UUID      string      `gorm:"type:varchar(40);primaryKey;not null"`
	ClubID    string      `gorm:"type:varchar(40);not null"`
	RankID    string      `gorm:"type:varchar(40);not null"`
	CreatedAt time.Time   `gorm:"->;type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
	Label     types.Label `gorm:"type:json;not null"`
	Name      string      `gorm:"type:varchar(40);not null"`
	Gender    int32       `gorm:"type:tinyint;not null"`
}

func (m Customer) TableName() string {
	return TableName()
}

func (m Customer) GetUUID() string {
	return m.UUID
}
