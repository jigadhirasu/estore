package fs

import (
	"estore/internal/mariadb"
	"estore/internal/models/filemap"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm"
)

func F1FindFilestore(dbname string) rex.Func1[string, string] {
	return func(ctx rex.Context, id string) (string, error) {
		db := rex.Get[*gorm.DB](ctx, mariadb.DBKey(dbname))

		filestore := filemap.FileStore{}
		if tx := db.First(&filestore, "id = ?", id); tx.Error != nil {
			return "", tx.Error
		}

		return filestore.FileName, nil
	}
}
