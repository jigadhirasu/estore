package transfer1

import (
	"estore/internal/models/club"
	"estore/protoc/clubpb"
)

func ClubPbToModel(a *clubpb.Club) club.Club {
	return club.Club{
		UUID: a.Id,
		Name: a.Name,
	}
}
