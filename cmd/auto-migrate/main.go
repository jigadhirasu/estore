package main

import (
	"context"
	"estore/internal/mariadb"
	"estore/internal/models/club"
	"estore/internal/models/filemap"

	"github.com/jigadhirasu/rex"
)

var autoMigrateTable = []mariadb.Table{
	&club.Club{},
	&filemap.FileMap{},
	filemap.FileStore{},
}

func main() {

	ctx := rex.NewContext(context.TODO())

	dbname := "test"

	pipe := rex.Pipe4(
		rex.Tap(mariadb.F0CreateDatabase[mariadb.Table](dbname)),
		rex.Once(mariadb.F0GetDB[mariadb.Table](dbname)),
		rex.Map[mariadb.Table, int64](mariadb.F1AutoMigrate(dbname)),
		rex.Reduce[int64, int64](func(a, b int64) int64 {
			return a + b
		}),
	)(
		rex.From[mariadb.Table](
			autoMigrateTable...,
		),
	)(ctx)

	result, err := pipe.ToSlice()
	if err != nil {
		panic(err)
	}

	if len(autoMigrateTable) != int(result[0]) {
		panic("auto migrate failed")
	}
}
