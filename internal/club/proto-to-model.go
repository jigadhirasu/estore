package club

import (
	"estore/internal/models/club"
	"estore/protoc/clubpb"

	"github.com/google/uuid"
	"github.com/jigadhirasu/rex"
)

func T1ProtoToModel(a *clubpb.Club) club.Club {

	m := club.Club{
		UUID: a.Id,
		Name: a.Name,
	}

	if m.UUID == "" {
		m.UUID = uuid.NewString()
	}

	return m
}
func F1ProtoToModel(ctx rex.Context, a *clubpb.Club) (club.Club, error) {
	return T1ProtoToModel(a), nil
}
