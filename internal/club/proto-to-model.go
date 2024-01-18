package club

import (
	"estore/internal/models/club"
	"estore/protoc/clubpb"

	"github.com/google/uuid"
	"github.com/jigadhirasu/rex"
)

func T1CloubProtoToModel(a *clubpb.Club) club.Club {
	return club.Club{
		UUID: uuid.NewString(),
		Name: a.Name,
	}
}
func F1CloubProtoToModel(ctx rex.Context, a *clubpb.Club) (club.Club, error) {
	return T1CloubProtoToModel(a), nil
}
