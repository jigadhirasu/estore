package mariadb

import (
	"fmt"

	"github.com/jigadhirasu/rex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func F0CreateDatabase[A Table](dbname string) rex.Func0[A] {
	return func(ctx rex.Context, a A) error {

		dsn := T1ParseDatabaseDSN(``)
		db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
		if err != nil {
			return err
		}

		raw := fmt.Sprintf(
			`
				CREATE DATABASE IF NOT EXISTS %s
				DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci
			`,
			Name(dbname),
		)

		tx := db.Exec(raw)
		if tx.Error != nil {
			return tx.Error
		}

		original, err1 := db.DB()
		err2 := original.Close()
		if err1 != nil || err2 != nil {
			return fmt.Errorf(`err1: %v, err2: %v`, err1, err2)

		}

		return nil
	}
}
