package club

import (
	"estore/internal/models/club"
	"estore/protoc/clubpb"

	"github.com/jigadhirasu/rex"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func T1ModelToProto(a club.Club) *clubpb.Club {
	return &clubpb.Club{
		Id:   a.UUID,
		Name: a.Name,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: a.CreatedAt.Unix(),
		},
	}
}
func F1ModelToProto(ctx rex.Context, a club.Club) (*clubpb.Club, error) {
	return T1ModelToProto(a), nil
}
