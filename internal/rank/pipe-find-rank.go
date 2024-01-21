package rank

import (
	"estore/internal/mariadb"
	"estore/internal/models/rank"
	"estore/internal/util"
	"estore/protoc/clubpb"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeFindClub(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.FindRankRequest, *clubpb.FindRankResponse] {
	return rex.Pipe5(
		rex.Map(T1BuildFindCondition),
		rex.MergeMap(mariadb.H1Find[rank.Rank](dbKey)),
		rex.Map(F1ModelToProto),
		rex.ReduceSlice[*clubpb.Rank](),
		rex.Map[[]*clubpb.Rank, *clubpb.FindRankResponse](func(ctx rex.Context, a []*clubpb.Rank) (*clubpb.FindRankResponse, error) {
			return &clubpb.FindRankResponse{
				Data: a,
			}, nil
		}),
	)
}

func T1BuildFindCondition(ctx rex.Context, req *clubpb.FindRankRequest) ([]clause.Expression, error) {

	expression := []clause.Expression{}

	if len(req.Id) > 0 {
		expression = append(expression, clause.IN{
			Column: "uuid",
			Values: array.Map[string, any](func(a string) any { return a })(req.Id),
		})
	}

	if len(req.ClubId) > 0 {
		expression = append(expression, clause.IN{
			Column: "club_id",
			Values: array.Map[string, any](func(a string) any { return a })(req.ClubId),
		})
	}

	if req.LikeName != "" {
		expression = append(expression, clause.Like{
			Column: "name",
			Value:  req.LikeName,
		})
	}

	if len(req.Priority) > 0 {
		expression = append(expression, clause.IN{
			Column: "priority",
			Values: array.Map[int32, any](func(a int32) any { return a })(req.Priority),
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
