package commodityspecrule

import "time"

func TableName() string {
	return "t_commodity_spec_rule"
}

// 開放給商店的規則
type CommoditySpecRule struct {
	UUID        string    `gorm:"type:varchar(40);primaryKey;not null"`
	Version     int       `gorm:"type:int;primaryKey;not null"`
	StoreID     string    `gorm:"type:varchar(40);not null"`
	Name        string    `gorm:"type:varchar(40);not null"`
	CommodityID string    `gorm:"type:varchar(40);not null"`
	Sku         string    `gorm:"type:varchar(40);not null"`
	Rule        int32     `gorm:"type:tinyint;not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m CommoditySpecRule) TableName() string {
	return TableName()
}

func (m CommoditySpecRule) GetUUID() string {
	return m.UUID
}
