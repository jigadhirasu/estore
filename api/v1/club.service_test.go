package v1

import (
	"context"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"
	"slices"
	"testing"

	"github.com/IBM/fp-go/array"
	"github.com/stretchr/testify/assert"
)

func testClubData() []*clubpb.Club {
	return []*clubpb.Club{
		{Name: "食品區"},
		{Name: "家電"},
		{Name: "3C"},
		{Name: "服飾"},
	}
}

func TestCreateClub(t *testing.T) {
	service := ClubService{}

	resp, err := service.CreateClub(context.TODO(), &clubpb.CreateClubRequest{
		Data: testClubData(),
	})

	assert.NoError(t, err)

	assert.Equal(t, int32(4), resp.Response.AffectedRows)
}

func TestFindClub(t *testing.T) {

	service := ClubService{}

	resp, err := service.FindClub(context.TODO(), &clubpb.FindClubRequest{
		Request: &typepb.Request{
			OrderBy: []*typepb.OrderBy{
				{Field: "name", Order: typepb.Order_DESC},
			},
		},
	})

	assert.NoError(t, err)

	answer := array.Map[*clubpb.Club, string](func(a *clubpb.Club) string { return a.Name })(testClubData())
	slices.SortFunc(answer, func(a, b string) int {
		if a < b {
			return 1
		}
		return -1
	})

	names := array.Map[*clubpb.Club, string](func(a *clubpb.Club) string { return a.Name })(resp.Data)

	assert.Equal(t, answer, names)

}
