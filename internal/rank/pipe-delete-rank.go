package rank

import (
	"estore/internal/mariadb"
	"estore/internal/models/rank"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeDeleteRank(dbKey string, clauses ...clause.Expression) rex.PipeLine[string, *clubpb.DeleteRankResponse] {
	return rex.Pipe5(
		rex.Map[string, rank.Rank](func(ctx rex.Context, a string) (rank.Rank, error) {
			return rank.Rank{UUID: a}, nil
		}),
		rex.ReduceSlice[rank.Rank](),
		rex.Map(mariadb.F1Delete[[]rank.Rank](dbKey, clauses...)),
		rex.Reduce[int64, int64](func(a, b int64) int64 {
			return a + b
		}),
		rex.Map[int64, *clubpb.DeleteRankResponse](func(ctx rex.Context, a int64) (*clubpb.DeleteRankResponse, error) {
			return &clubpb.DeleteRankResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
