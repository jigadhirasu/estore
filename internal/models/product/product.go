package product

import "estore/internal/types"

func TableName() string {
	return "t_product"
}

// TODO 還沒完成之後再設計
type Product struct {
	UUID        string          `gorm:"type:varchar(40);primaryKey;not null"`
	Version     int             `gorm:"type:int;primaryKey;not null"`
	Name        string          `gorm:"type:varchar(40);not null"`
	State       int32           `gorm:"type:tinyint;not null"`
	Price       types.Price     `gorm:"not null"`
	Url         string          `gorm:"type:varchar(40);not null"`
	Brief       string          `gorm:"type:varchar(63);not null"`
	Description string          `gorm:"type:varchar(255);not null"`
	Label       types.Label     `gorm:"type:json;not null"`
	SEO         types.SEO       `gorm:"type:json;not null"`
	OpenGraph   types.OpenGraph `gorm:"type:json;not null"`
	CreatedAt   string          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Product) TableName() string {
	return TableName()
}

func (m Product) GetUUID() string {
	return m.UUID
}
