package v1

import (
	"context"
	"estore/protoc/clubpb"
	"estore/protoc/typepb"
	"slices"
	"strings"
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

func TestUpdateClub(t *testing.T) {

	service := ClubService{}

	resp1, err := service.FindClub(context.TODO(), &clubpb.FindClubRequest{
		Request: &typepb.Request{
			OrderBy: []*typepb.OrderBy{
				{Field: "name", Order: typepb.Order_DESC},
			},
		},
	})

	assert.NoError(t, err)

	assert.Equal(t, 4, len(resp1.Data))

	for i, cl := range resp1.Data {
		if i == 2 {
			break
		}

		cl.Name = cl.Name + "1"
	}

	resp2, err := service.UpdateClub(context.TODO(), &clubpb.UpdateClubRequest{
		Data: resp1.Data,
	})

	assert.NoError(t, err)

	assert.Equal(t, int32(2), resp2.Response.AffectedRows)

	resp3, err := service.FindClub(context.TODO(), &clubpb.FindClubRequest{
		Request: &typepb.Request{
			OrderBy: []*typepb.OrderBy{
				{Field: "name", Order: typepb.Order_DESC},
			},
		},
	})

	assert.NoError(t, err)

	answer := array.Map[*clubpb.Club, string](func(a *clubpb.Club) string { return a.Name })(resp1.Data)
	slices.SortFunc(answer, func(a, b string) int {
		if a < b {
			return 1
		}
		return -1
	})

	names := array.Map[*clubpb.Club, string](func(a *clubpb.Club) string { return a.Name })(resp3.Data)

	assert.Equal(t, answer, names)

	for _, cl := range resp1.Data {
		cl.Name = strings.ReplaceAll(cl.Name, "1", "")
	}

	resp4, err := service.UpdateClub(context.TODO(), &clubpb.UpdateClubRequest{
		Data: resp1.Data,
	})

	assert.NoError(t, err)

	assert.Equal(t, int32(2), resp4.Response.AffectedRows)
}
