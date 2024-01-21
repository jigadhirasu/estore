package fs

import (
	"estore/internal/mariadb"
	"estore/internal/models/filemap"
	"estore/protoc/fspb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeCreateFile(dbKey string, clauses ...clause.Expression) rex.PipeLine[*fspb.Dir, *fspb.CreateOrUpdateDirFileResponse] {
	return rex.Pipe7(
		rex.Map(F1ResourceRoot),
		rex.MergeMap[*fspb.Dir, *fspb.File](H1ExpandCreateFile),
		rex.Map(F1MkdirAndSaveFile),
		rex.Map(F1FileToFilestore),
		rex.ReduceSlice[filemap.FileStore](),
		rex.Map(mariadb.F1Insert[[]filemap.FileStore](dbKey, clauses...)),
		rex.Map[int64, *fspb.CreateOrUpdateDirFileResponse](func(ctx rex.Context, a int64) (*fspb.CreateOrUpdateDirFileResponse, error) {
			return &fspb.CreateOrUpdateDirFileResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
