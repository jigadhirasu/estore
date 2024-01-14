package storehouse

import (
	"estore/internal/types"
	"time"
)

func TableName() string {
	return "t_storehouse"
}

type Storehouse struct {
	UUID      string        `gorm:"type:varchar(40);primaryKey;not null"`
	Version   int           `gorm:"type:int;primaryKey;not null"`
	Name      string        `gorm:"type:varchar(40);not null"`
	Phone     types.Phone   `gorm:"not null"`
	Address   types.Address `gorm:"not null"`
	CreatedAt time.Time     `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Storehouse) TableName() string {
	return TableName()
}

func (m Storehouse) GetUUID() string {
	return m.UUID
}
