package mariadb

import (
	"github.com/jigadhirasu/rex"
	"gorm.io/gorm"
)

func Begin(ctx rex.Context, dbname string) (*gorm.DB, error) {
	db := rex.Get[*gorm.DB](ctx, DBKey(dbname))
	tx := db.Begin()

	rex.Set(ctx, TxKey(dbname), tx)
	return tx, nil
}
func F0Begin[A any](dbname string) rex.Func0[A] {
	return func(ctx rex.Context, a A) error {
		_, err := Begin(ctx, dbname)
		return err
	}
}
func F1Begin[A any](dbname string) rex.Func1[A, *gorm.DB] {
	return func(ctx rex.Context, a A) (*gorm.DB, error) {
		return Begin(ctx, dbname)
	}
}

func Commit(ctx rex.Context, dbname string) error {
	tx, ok := rex.Maybe[*gorm.DB](ctx, TxKey(dbname))
	if !ok {
		return nil
	}

	rex.Delete(ctx, TxKey(dbname))
	return tx.Commit().Error
}
func F0Commit[A any](dbname string) rex.Func0[A] {
	return func(ctx rex.Context, a A) error {
		return Commit(ctx, dbname)
	}
}

func Rollback(ctx rex.Context, dbname string) error {
	tx, ok := rex.Maybe[*gorm.DB](ctx, TxKey(dbname))
	if !ok {
		return nil
	}

	rex.Delete(ctx, TxKey(dbname))
	return tx.Rollback().Error
}
func F0Rollback[A any](dbname string) rex.Func0[A] {
	return func(ctx rex.Context, a A) error {
		return Rollback(ctx, dbname)
	}
}
