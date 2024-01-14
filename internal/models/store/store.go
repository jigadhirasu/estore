package store

func TableName() string {
	return "t_store"
}

type Store struct {
	UUID       string    `gorm:"type:varchar(40);primaryKey;not null"`
	Version    int       `gorm:"type:int;primaryKey;not null"`
	Type       int32     `gorm:"type:tinyint;not null"`
	State      int32     `gorm:"type:tinyint;not null"`
	ClubID     string    `gorm:"type:varchar(40);not null"`
	Storehouse string    `gorm:"type:varchar(40);not null"`
	Currency   int32     `gorm:"type:tinyint;not null"`
	CreatedAt  string    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
	Info       StoreInfo `gorm:"not null"`
}

func (m Store) TableName() string {
	return TableName()
}

func (m Store) GetUUID() string {
	return m.UUID
}
