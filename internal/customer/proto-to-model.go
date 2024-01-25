package customer

import (
	"estore/internal/models/customer"
	"estore/protoc/clubpb"

	"github.com/google/uuid"
	"github.com/jigadhirasu/rex"
)

func T1ProtoToModel(a *clubpb.Customer) customer.Customer {

	m := customer.Customer{
		UUID:      a.Id,
		ClubID:    a.ClubId,
		RankID:    a.RankId,
		Label:     a.Label,
		CreatedAt: a.CreatedAt.AsTime(),
		Name:      a.Name,
		Gender:    int32(a.Gender),
	}

	if m.UUID == "" {
		m.UUID = uuid.NewString()
	}

	return m
}
func F1ProtoToModel(ctx rex.Context, a *clubpb.Customer) (customer.Customer, error) {
	return T1ProtoToModel(a), nil
}
