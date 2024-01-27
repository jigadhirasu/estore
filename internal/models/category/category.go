package category

func TableName() string {
	return "t_category"
}

type Category struct {
	UUID       string `gorm:"type:varchar(40);primaryKey;not null"`
	Version    int    `gorm:"type:int;primaryKey;not null"`
	Name       string `gorm:"type:varchar(40);not null"`
	Path       string `gorm:"type:varchar(40);index:parent_id;not null"`
	BusinessID string `gorm:"type:varchar(40);index:business_id;not null"`
	ClubID     string `gorm:"type:varchar(40);index:club_id;not null"`
	StoreID    string `gorm:"type:varchar(40);index:store_id;not null"`
	CreatedAt  string `gorm:"->;type:timestamp;default:CURRENT_TIMESTAMP;index:created_at;not null"`
}

func (m Category) TableName() string {
	return TableName()
}

func (m Category) GetUUID() string {
	return m.UUID
}
