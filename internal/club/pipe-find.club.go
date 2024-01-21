package club

import (
	"estore/internal/mariadb"
	"estore/internal/models/club"
	"estore/internal/util"
	"estore/protoc/clubpb"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeFindClub(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.FindClubRequest, *clubpb.FindClubResponse] {
	return rex.Pipe5(
		rex.Map(T1BuildFindCondition),
		rex.MergeMap(mariadb.H1Find[club.Club](dbKey)),
		rex.Map(F1ModelToProto),
		rex.ReduceSlice[*clubpb.Club](),
		rex.Map[[]*clubpb.Club, *clubpb.FindClubResponse](func(ctx rex.Context, a []*clubpb.Club) (*clubpb.FindClubResponse, error) {
			return &clubpb.FindClubResponse{
				Data: a,
			}, nil
		}),
	)
}

func T1BuildFindCondition(ctx rex.Context, req *clubpb.FindClubRequest) ([]clause.Expression, error) {

	expression := []clause.Expression{}

	if len(req.Id) > 0 {
		expression = append(expression, clause.IN{
			Column: "uuid",
			Values: array.Map[string, any](func(a string) any { return a })(req.Id),
		})
	}

	if req.LikeName != "" {
		expression = append(expression, clause.Like{
			Column: "name",
			Value:  req.LikeName,
		})
	}

	expression = append(expression,
		util.LimitOffset(req.Request),
		util.OrderBy(clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{Column: clause.Column{Name: "created_at"}, Desc: true},
			},
		})(req.Request),
	)

	return expression, nil
}
