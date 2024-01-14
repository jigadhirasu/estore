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

type Range struct {
	Min int32
	Max int32
}

type Ranges []Range

func (a *Ranges) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

func (a Ranges) Value() (driver.Value, error) {
	if !reflect.ValueOf(a).IsValid() {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a Ranges) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if len(a) == 0 {
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

func (a Ranges) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return JsonGormDBDataType(db, field)
}
