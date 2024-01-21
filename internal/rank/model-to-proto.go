package rank

import (
	"estore/internal/models/rank"
	"estore/internal/types"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/IBM/fp-go/array"
	"github.com/jigadhirasu/rex"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func T1ModelToProto(a rank.Rank) *clubpb.Rank {
	return &clubpb.Rank{
		Id:       a.UUID,
		ClubId:   a.Club,
		Name:     a.Name,
		Priority: a.Priority,
		Condition: array.Map(func(a types.Condition) *typepb.Condition {
			return &typepb.Condition{
				CumulativePrice: a.CumulativePrice,
				OneTimePrice:    a.OneTimePrice,
				Times:           a.Times,
			}
		})(a.Conditions),
		CreatedAt: timestamppb.New(a.CreatedAt),
	}
}
func F1ModelToProto(ctx rex.Context, a rank.Rank) (*clubpb.Rank, error) {
	return T1ModelToProto(a), nil
}
