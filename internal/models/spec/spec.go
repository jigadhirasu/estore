package spec

import (
	"estore/internal/types"
	"time"
)

func TableName() string {
	return "t_spec"
}

type Spec struct {
	UUID        string      `gorm:"type:varchar(40);primaryKey;not null"`
	Version     int         `gorm:"type:int;primaryKey;not null"`
	CommodityID string      `gorm:"type:varchar(40);not null"`
	Volume      int32       `gorm:"type:smallint;not null"`
	Weight      int32       `gorm:"type:smallint;not null"`
	Label       types.Label `gorm:"type:json;not null"`
	CreatedAt   time.Time   `gorm:"->;type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Spec) TableName() string {
	return TableName()
}

func (m Spec) GetUUID() string {
	return m.UUID
}
