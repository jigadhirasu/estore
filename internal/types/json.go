package types

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func JsonGormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "json"
	case "mysql":
		return "json"
	case "postgres":
		return "jsonb"
	}
	return ""
}
