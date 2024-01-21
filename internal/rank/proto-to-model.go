package rank

import (
	"estore/internal/models/rank"
	"estore/internal/types"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"

	"github.com/IBM/fp-go/array"
	"github.com/google/uuid"
	"github.com/jigadhirasu/rex"
)

func T1ProtoToModel(a *clubpb.Rank) rank.Rank {

	m := rank.Rank{
		UUID:     a.Id,
		Club:     a.ClubId,
		Name:     a.Name,
		Priority: a.Priority,
		Conditions: array.Map(func(a *typepb.Condition) types.Condition {
			return types.Condition{
				CumulativePrice: a.CumulativePrice,
				OneTimePrice:    a.OneTimePrice,
				Times:           a.Times,
			}
		})(a.Condition),
	}

	if m.UUID == "" {
		m.UUID = uuid.NewString()
	}

	return m
}
func F1ProtoToModel(ctx rex.Context, a *clubpb.Rank) (rank.Rank, error) {
	return T1ProtoToModel(a), nil
}
