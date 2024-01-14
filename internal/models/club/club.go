package club

import (
	"time"
)

// 用view拉出每個ID的最新紀錄
func TableName() string {
	return "t_club"
}

type Club struct {
	UUID      string    `gorm:"type:varchar(40);primaryKey;not null"`
	Name      string    `gorm:"type:varchar(40);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Club) TableName() string {
	return TableName()
}

func (m Club) GetUUID() string {
	return m.UUID
}
