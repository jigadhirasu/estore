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

type Condition struct {
	CumulativePrice int32 // 累計金額
	OneTimePrice    int32 // 一次性的金額
	Times           int32 // 次數
}

type Conditions []Condition

func (a *Conditions) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

func (a Conditions) Value() (driver.Value, error) {
	if !reflect.ValueOf(a).IsValid() {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a Conditions) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (a Conditions) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return JsonGormDBDataType(db, field)
}
