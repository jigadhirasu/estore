package club

import (
	"estore/internal/mariadb"
	"estore/internal/models/club"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeUpdateClub(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Club, *clubpb.UpdateClubResponse] {
	return rex.Pipe6(
		rex.Tap[*clubpb.Club](F0CheckIdIsValid),
		rex.Tap[*clubpb.Club](F0CheckNameIsValid),
		rex.Map(F1ProtoToModel),
		rex.Map(mariadb.F1Update[club.Club](dbKey, clauses...)),
		rex.Reduce[int64, int64](func(a, b int64) int64 {
			return a + b
		}),
		rex.Map[int64, *clubpb.UpdateClubResponse](func(ctx rex.Context, a int64) (*clubpb.UpdateClubResponse, error) {
			return &clubpb.UpdateClubResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
