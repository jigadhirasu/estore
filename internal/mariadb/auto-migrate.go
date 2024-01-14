package mariadb

import (
	"github.com/jigadhirasu/rex"
	"gorm.io/gorm"
)

type Table interface {
	TableName() string
}

// 建立資料庫
func F1AutoMigrate(dbname string) rex.Func1[Table, int64] {
	return func(ctx rex.Context, table Table) (int64, error) {

		db := rex.Get[*gorm.DB](ctx, DBKey(dbname))

		if err := db.AutoMigrate(table); err != nil {
			return 0, err
		}

		return 1, nil
	}
}
