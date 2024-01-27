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

func PipeUpdateCustomer(dbKey string, clauses ...clause.Expression) rex.PipeLine[*clubpb.Customer, *clubpb.UpdateCustomerResponse] {
	return rex.Pipe7(
		rex.Tap[*clubpb.Customer](validate.IsIdValid),
		rex.Tap[*clubpb.Customer](validate.IsNameValid),
		rex.Tap[*clubpb.Customer](validate.IsRankIdValid),
		rex.Map(F1ProtoToModel),
		rex.Map(mariadb.F1Update[customer.Customer](dbKey, clauses...)),
		rex.Reduce[int64, int64](func(a, b int64) int64 {
			return a + b
		}),
		rex.Map(func(ctx rex.Context, a int64) (*clubpb.UpdateCustomerResponse, error) {
			return &clubpb.UpdateCustomerResponse{
				Response: &typepb.Response{
					AffectedRows: int32(a),
				},
			}, nil
		}),
	)
}
