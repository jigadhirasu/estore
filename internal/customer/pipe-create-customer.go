package customer

import (
	"estore/internal/mariadb"
	"estore/internal/models/customer"
	validate "estore/internal/validdate"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/jigadhirasu/rex"
	"gorm.io/gorm/clause"
)

func PipeCreateCustomer(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Customer, *clubpb.CreateCustomerResponse] {
	return rex.Pipe7(
		rex.Tap[*clubpb.Customer](validate.IsNameValid),
		rex.Tap[*clubpb.Customer](validate.IsClubIdValid),
		rex.Tap[*clubpb.Customer](validate.IsRankIdValid),
		rex.Map(F1ProtoToModel),
		rex.ReduceSlice[customer.Customer](),
		rex.Map(mariadb.F1Insert[[]customer.Customer](dbKey, clauses...)),
		rex.Map[int64, *clubpb.CreateCustomerResponse](func(ctx rex.Context, a int64) (*clubpb.CreateCustomerResponse, error) {
			return &clubpb.CreateCustomerResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
