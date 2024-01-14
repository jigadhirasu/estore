package filemap

func TableName(tp Type) string {
	return "t_file" + string(tp)
}

type Type string

const (
	Store Type = "store"
	Map   Type = "map"
)

type FileMap struct {
	ID         int64
	FileID     string `gorm:"type:varchar(40);uniqueIndex:relation_file;not null"`
	RelationID string `gorm:"type:varchar(40);uniqueIndex:relation_file;not null"`
}

func (m FileMap) TableName() string {
	return TableName(Map)
}
