package rank

import (
	"estore/internal/mariadb"
	"estore/internal/models/rank"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeUpdateRank(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Rank, *clubpb.UpdateRankResponse] {
	return rex.Pipe7(
		rex.Tap[*clubpb.Rank](F0CheckIdIsValid),
		rex.Tap[*clubpb.Rank](F0CheckIdIsValid),
		rex.Tap[*clubpb.Rank](F0CheckNameIsValid),
		rex.Map(F1ProtoToModel),
		rex.Map(mariadb.F1Update[rank.Rank](dbKey, clauses...)),
		rex.Reduce[int64, int64](func(a, b int64) int64 {
			return a + b
		}),
		rex.Map[int64, *clubpb.UpdateRankResponse](func(ctx rex.Context, a int64) (*clubpb.UpdateRankResponse, error) {
			return &clubpb.UpdateRankResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
