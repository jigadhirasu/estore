package mariadb

import (
	"estore/pkg/module/logger"

	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func F1Delete[A any](dbKey string, clauses ...clause.Expression) rex.Func1[A, int64] {
	return func(ctx rex.Context, data A) (int64, error) {

		db := rex.Get[*gorm.DB](ctx, dbKey)

		tx := db.Clauses(clauses...).Delete(data)

		if tx.Error != nil {
			logger.Out(ctx).Error("DeleteToDB",
				zap.Error(tx.Error),
				zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
					return tx.Clauses(clauses...).Updates(data)
				})),
			)
			return 0, tx.Error
		}

		logger.Out(ctx).Info("DeleteToDB",
			logger.JSON("data", data),
		)

		return tx.RowsAffected, nil
	}
}
