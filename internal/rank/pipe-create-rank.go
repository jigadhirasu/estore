package rank

import (
	"estore/internal/mariadb"
	"estore/internal/models/rank"
	validate "estore/internal/validdate"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeCreateRank(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Rank, *clubpb.CreateRankResponse] {
	return rex.Pipe6(
		rex.Tap[*clubpb.Rank](validate.IsNameValid),
		rex.Tap[*clubpb.Rank](IsConditionValid),
		rex.Map(F1ProtoToModel),
		rex.ReduceSlice[rank.Rank](),
		rex.Map(mariadb.F1Insert[[]rank.Rank](dbKey, clauses...)),
		rex.Map(func(ctx rex.Context, a int64) (*clubpb.CreateRankResponse, error) {
			return &clubpb.CreateRankResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
