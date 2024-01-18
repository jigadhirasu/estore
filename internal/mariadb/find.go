package mariadb

import (
	"estore/pkg/module/logger"

	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func F1Find[A Table](dbKey string) rex.Func1[[]clause.Expression, []A] {
	return func(ctx rex.Context, clauses []clause.Expression) ([]A, error) {

		db := rex.Get[*gorm.DB](ctx, dbKey)

		var data []A
		tx := db.Clauses(clauses...).Find(&data)

		if tx.Error != nil {
			logger.Out(ctx).Error("FindFromDB",
				zap.Error(tx.Error),
				zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
					return tx.Clauses(clauses...).Find(&data)
				})),
			)
			return data, tx.Error
		}

		return data, nil
	}
}
func H1Find[A Table](dbKey string) rex.HFunc1[[]clause.Expression, A] {
	return rex.HFunc1[[]clause.Expression, A](func(ctx rex.Context, clauses []clause.Expression) rex.Iterable[A] {

		db := rex.Get[*gorm.DB](ctx, dbKey)

		rows, err := db.Model(new(A)).Clauses(clauses...).Rows()
		if err != nil {
			logger.Out(ctx).Error("FindStreamFromDB",
				zap.Error(err),
				zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
					return tx.Clauses(clauses...).Find(&[]A{})
				})),
			)

			return rex.FromItem[A](rex.ItemError[A](err))
		}

		ch := make(chan rex.Item[A])

		go func() {
			defer close(ch)
			defer rows.Close()

			var a A
			for rows.Next() {
				if err := db.ScanRows(rows, &a); err != nil {
					logger.Out(ctx).Error("FindFromDB",
						zap.Error(err),
						zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
							return tx.Clauses(clauses...).Find(&a)
						})),
					)
					ch <- rex.ItemError[A](err)
					return
				}

				ch <- rex.ItemOf[A](a)
			}
		}()

		return rex.FromChanItem[A](ch)
	})
}
