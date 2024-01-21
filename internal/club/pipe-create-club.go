package club

import (
	"estore/internal/mariadb"
	"estore/internal/models/club"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeCreateClub(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Club, *clubpb.CreateClubResponse] {
	return rex.Pipe5(
		rex.Tap[*clubpb.Club](F0CheckNameIsValid),
		rex.Map(F1ProtoToModel),
		rex.ReduceSlice[club.Club](),
		rex.Map(mariadb.F1Insert[[]club.Club](dbKey, clauses...)),
		rex.Map[int64, *clubpb.CreateClubResponse](func(ctx rex.Context, a int64) (*clubpb.CreateClubResponse, error) {
			return &clubpb.CreateClubResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
