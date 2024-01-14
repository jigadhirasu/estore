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

type Array[A any] []A

func (a *Array[any]) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

func (a Array[any]) Value() (driver.Value, error) {
	if !reflect.ValueOf(a).IsValid() {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a Array[any]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if !reflect.ValueOf(a).IsValid() {
		return gorm.Expr("?", "[]")
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

// GormDBDataType gorm db data type
func (a *Array[any]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return JsonGormDBDataType(db, field)
}
