package customer

import (
	"estore/internal/mariadb"
	"estore/internal/models/customer"
	"estore/internal/util"
	"estore/protoc/clubpb"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeFindCustomer(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.FindCustomerRequest, *clubpb.FindCustomerResponse] {
	return rex.Pipe5(
		rex.Map(T1BuildFindCondition),
		rex.MergeMap(mariadb.H1Find[customer.Customer](dbKey)),
		rex.Map(F1ModelToProto),
		rex.ReduceSlice[*clubpb.Customer](),
		rex.Map(func(ctx rex.Context, a []*clubpb.Customer) (*clubpb.FindCustomerResponse, error) {
			return &clubpb.FindCustomerResponse{
				Data: a,
			}, nil
		}),
	)
}

func T1BuildFindCondition(ctx rex.Context, req *clubpb.FindCustomerRequest) ([]clause.Expression, error) {

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
