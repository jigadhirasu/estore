package types

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Email struct {
	Account string
	Domain  string
}

func (a *Email) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

func (a Email) Value() (driver.Value, error) {
	if !reflect.ValueOf(a).IsValid() {
		return []byte("{}"), nil
	}

	return json.Marshal(a)
}

func (a Email) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if !reflect.ValueOf(a).IsValid() {
		return gorm.Expr("?", "{}")
	}

	data, _ := json.Marshal(a)

	switch db.Dialector.Name() {
	case "mysql":
		if v, ok := db.Dialector.(*mysql.Dialector); ok && !strings.Contains(v.ServerVersion, "MariaDB") {
			return gorm.Expr("CAST(? AS JSON)", string(data))
		}
	}

	return gorm.Expr("?", string(data))
}

func (a Email) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return JsonGormDBDataType(db, field)
}
