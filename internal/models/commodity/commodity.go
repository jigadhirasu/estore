package commodity

import (
	"estore/internal/types"
	"time"
)

func TableName() string {
	return "t_commodity"
}

type Commodity struct {
	UUID      string      `gorm:"type:varchar(40);primaryKey;not null"`
	Version   int         `gorm:"type:int;primaryKey;not null"`
	ProductID string      `gorm:"type:varchar(40);not null"`
	State     int32       `gorm:"type:tinyint;not null"`
	Cost      float64     `gorm:"type:double;not null"`
	Label     types.Label `gorm:"type:json;not null"`
	CreatedAt time.Time   `gorm:"->;type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Commodity) TableName() string {
	return TableName()
}

func (m Commodity) GetUUID() string {
	return m.UUID
}
