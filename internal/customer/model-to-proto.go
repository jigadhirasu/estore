package customer

import (
	"estore/internal/models/customer"
	"estore/protoc/clubpb"

	"github.com/jigadhirasu/rex"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func T1ModelToProto(a customer.Customer) *clubpb.Customer {
	return &clubpb.Customer{
		Id:     a.UUID,
		ClubId: a.ClubID,
		RankId: a.RankID,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: a.CreatedAt.Unix(),
		},
		Label:  a.Label,
		Name:   a.Name,
		Gender: clubpb.Gender(a.Gender),
	}
}
func F1ModelToProto(ctx rex.Context, a customer.Customer) (*clubpb.Customer, error) {
	return T1ModelToProto(a), nil
}
