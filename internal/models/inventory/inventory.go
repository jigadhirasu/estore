package inventory

import (
	"time"
)

func TableName() string {
	return "t_inventory"
}

type Inventory struct {
	UUID      string    `gorm:"type:varchar(40);primaryKey;not null"`
	Version   int       `gorm:"type:int;primaryKey;not null"`
	OrderID   string    `gorm:"type:varchar(40);not null"`
	Sku       string    `gorm:"type:varchar(127);not null"`
	Quantity  int32     `gorm:"type:tinyint;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Inventory) TableName() string {
	return TableName()
}

func (m Inventory) GetUUID() string {
	return m.UUID
}
