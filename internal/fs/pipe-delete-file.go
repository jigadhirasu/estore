package fs

import (
	"estore/protoc/fspb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
)

func PipeDeleteFile(dbKey string) rex.PipeLine[*fspb.Dir, *fspb.DeleteDirFileResponse] {
	return rex.Pipe5(
		rex.Map(F1ResourceRoot),
		rex.MergeMap(H1ExpandDeleteDir),
		rex.Map(F1DeleteFileAndRelation(dbKey)),
		rex.Reduce[int64, int64](func(a, b int64) int64 {
			return a + b
		}),
		rex.Map[int64, *fspb.DeleteDirFileResponse](func(ctx rex.Context, a int64) (*fspb.DeleteDirFileResponse, error) {
			return &fspb.DeleteDirFileResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
