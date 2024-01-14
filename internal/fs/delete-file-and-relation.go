package fs

import (
	"estore/internal/models/filemap"
	"estore/pkg/module/logger"
	"estore/protoc/fspb"
	"os"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 使用獨立的transaction
func F1DeleteFileAndRelation(dbKey string) func(ctx rex.Context, dir *fspb.Dir) (int64, error) {
	return func(ctx rex.Context, dir *fspb.Dir) (int64, error) {

		db := rex.Get[*gorm.DB](ctx, dbKey)

		// 刪除整個目錄
		if len(dir.Files) < 1 {
			tx1 := db.Where(`file_id IN (SELECT id FROM t_filestore WHERE INSTR(file_name, ?) = 1)`, dir.Name).Delete(&filemap.FileMap{})
			if tx1.Error != nil {
				logger.Out(ctx).Error("DeleteFileAndRelation",
					zap.Error(tx1.Error),
					zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
						return tx.Where(`file_id IN (SELECT id FROM t_filestore WHERE INSTR(file_name, ?) = 1)`, dir.Name).Delete(&filemap.FileMap{})
					})),
				)
				return 0, tx1.Error
			}

			tx2 := db.Where(`INSTR(file_name, ?) = 1`, dir.Name).Find(&filemap.FileStore{})
			if tx2.Error != nil {
				logger.Out(ctx).Error("DeleteFileAndRelation",
					zap.Error(tx2.Error),
					zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
						return tx.Where(`INSTR(file_name, ?) = 1`, dir.Name).Delete(&filemap.FileStore{})
					})),
				)
				return 0, tx2.Error
			}

			if err := os.RemoveAll(dir.Name); err != nil {
				logger.Out(ctx).Error("DeleteFileAndRelation", zap.Error(err))
				return 0, err
			}

			return tx1.RowsAffected + tx2.RowsAffected, nil
		}

		// 刪除某些檔案
		filenames := []string{}
		for _, file := range dir.Files {
			name := dir.Name + "/" + file.Name
			filenames = append(filenames, name)
		}

		// 刪除關聯
		tx1 := db.Where(`file_id IN (SELECT id FROM t_filestore WHERE file_name IN ?)`, filenames).Delete(&filemap.FileMap{})
		if tx1.Error != nil {
			logger.Out(ctx).Error("DeleteFileAndRelation",
				zap.Error(tx1.Error),
				zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
					return tx.Where(`file_id IN (SELECT id FROM t_filestore WHERE file_name IN ?)`, filenames).Delete(&filemap.FileMap{})
				})),
			)
			return 0, tx1.Error
		}
		// 取出要刪除的檔案名稱
		removeFiles := []filemap.FileStore{}
		for _, filename := range filenames {
			fileStore := filemap.FileStore{}
			if tx := db.Where("INSTR(file_name, ?) = 1", filename).First(&fileStore); tx.Error != nil {
				logger.Out(ctx).Error("DeleteFileAndRelation",
					zap.Error(tx.Error),
					zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
						return tx.Where(`file_name IN ?`, filenames).Find(&removeFiles)
					})),
				)
			}
			removeFiles = append(removeFiles, fileStore)
		}
		// 刪除檔案名稱
		tx2 := db.Where(`file_name IN ?`,
			array.Map(func(a filemap.FileStore) string { return a.FileName })(removeFiles),
		).Delete(&filemap.FileStore{})
		if tx2.Error != nil {
			logger.Out(ctx).Error("DeleteFileAndRelation",
				zap.Error(tx2.Error),
				zap.String("sql", db.ToSQL(func(tx *gorm.DB) *gorm.DB {
					return tx.Where(`file_name IN ?`, filenames).Delete(&filemap.FileStore{})
				})),
			)
			return 0, tx2.Error
		}

		for _, file := range removeFiles {
			if err := os.Remove(file.FileName); err != nil {
				logger.Out(ctx).Error("DeleteFileAndRelation", zap.Error(err))
				return 0, err
			}
		}

		return tx1.RowsAffected + tx2.RowsAffected, nil
	}
}
