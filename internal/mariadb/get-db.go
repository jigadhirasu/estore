package mariadb

import (
	"os"
	"strconv"
	"time"

	"github.com/jigadhirasu/rex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(ctx rex.Context, dbname string) (*gorm.DB, error) {
	dsn := T1ParseDatabaseDSN(dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	maxidle, _ := strconv.Atoi(os.Getenv("MARIADB_MAX_IDLE"))
	if maxidle < 1 {
		maxidle = 4
	}

	maxopen, _ := strconv.Atoi(os.Getenv("MARIADB_MAX_OPEN"))
	if maxopen < 1 {
		maxopen = 16
	}

	originDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	originDB.SetMaxIdleConns(maxidle)
	originDB.SetMaxOpenConns(maxopen)
	originDB.SetConnMaxLifetime(time.Hour)

	rex.Set(ctx, DBKey(dbname), db)
	return db, nil
}
func F0GetDB[A any](dbname string) rex.Func0[A] {
	return func(ctx rex.Context, a A) error {
		_, err := GetDB(ctx, dbname)
		return err
	}
}
func F1GetDB[A any](dbname string) rex.Func1[A, *gorm.DB] {
	return func(ctx rex.Context, a A) (*gorm.DB, error) {
		return GetDB(ctx, dbname)
	}
}
