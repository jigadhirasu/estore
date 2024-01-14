package main

import (
	"context"
	"estore/configs"
	"estore/internal/mariadb"

	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jigadhirasu/rex"
)

func init() {
	configs.Develop()
}

func main() {
	// github://jigadhirasu:{UserAccessToken}@jigadhirasu/estore_migrate
	folder := os.Getenv("MIGRATE_FOLDER")

	// mysql://estore:q;4tj/6y7@tcp(localhost:14000)/?multiStatements=true

	db, err := mariadb.GetDB(rex.NewContext(context.TODO()), "test")
	if err != nil {
		panic(err)
	}

	db.Exec(`SET tidb_skip_isolation_level_check=1`)

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	driver, _ := mysql.WithInstance(sqlDB, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		folder,
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		panic(err)
	}

}
