package club

import (
	"errors"
	"estore/internal/mariadb"
	"estore/internal/models/club"
	"estore/internal/util"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeCreateClub(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Club, *clubpb.CreateClubResponse] {
	return rex.Pipe4(
		rex.Tap[*clubpb.Club](func(ctx rex.Context, a *clubpb.Club) error {
			if len(a.Name) < 1 {
				return errors.New("club name is empty")
			}

			return nil
		}),
		rex.Reduce(util.ReduceSliceFunc(T1CloubProtoToModel)),
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
