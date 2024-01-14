package store

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"estore/internal/types"
	"reflect"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type StoreInfo struct {
	CopyRight     string        `json:"copy_right"`
	Phone         types.Phone   `json:"phone"`
	Email         types.Email   `json:"email"`
	Address       types.Address `json:"address"`
	BusinessHours string        `json:"business_hours"`
	MaxCount      int32         `json:"max_count"`
	AllowCancel   bool          `json:"allow_cancel"`
	AllowReturn   bool          `json:"allow_return"`
	DisplayStock  bool          `json:"display_stock"`
}

func (a *StoreInfo) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

func (a StoreInfo) Value() (driver.Value, error) {
	if !reflect.ValueOf(a).IsValid() {
		return []byte("{}"), nil
	}

	return json.Marshal(a)
}

func (a StoreInfo) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

// GormDBDataType gorm db data type
func (a StoreInfo) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return types.JsonGormDBDataType(db, field)
}
