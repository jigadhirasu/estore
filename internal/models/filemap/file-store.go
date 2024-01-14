package filemap

// 紀錄uuid相對應的檔案位置
type FileStore struct {
	ID       string `gorm:"type:varchar(40);primaryKey;not null"`
	FileName string `gorm:"type:varchar(120);unique;not null"`
}

func (m FileStore) TableName() string {
	return TableName(Store)
}

func (m FileStore) GetUUID() string {
	return m.ID
}
