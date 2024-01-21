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

func testRankData(clubID string) []*clubpb.Rank {
	return []*clubpb.Rank{
		{ClubId: clubID, Name: "Gold", Priority: 1, Condition: []*typepb.Condition{
			{CumulativePrice: 10000, OneTimePrice: 1000, Times: 1},
		}},
		{ClubId: clubID, Name: "VIP", Priority: 2, Condition: []*typepb.Condition{
			{CumulativePrice: 20000, OneTimePrice: 2000, Times: 2},
		}},
		{ClubId: clubID, Name: "VVIP", Priority: 3, Condition: []*typepb.Condition{
			{CumulativePrice: 30000, OneTimePrice: 3000, Times: 3},
		}},
		{ClubId: clubID, Name: "VVVIP", Priority: 4, Condition: []*typepb.Condition{
			{CumulativePrice: 40000, OneTimePrice: 4000, Times: 4},
		}},
	}
}

func TestRankCreateFindUpdateDelete(t *testing.T) {

	service := ClubService{}

	resp1, _ := service.FindClub(context.TODO(), &clubpb.FindClubRequest{
		Request: &typepb.Request{
			OrderBy: []*typepb.OrderBy{
				{Field: "name", Order: typepb.Order_DESC},
			},
		},
	})

	resp2, err := service.CreateRank(context.TODO(), &clubpb.CreateRankRequest{
		Data: testRankData(resp1.Data[0].Id),
	})

	assert.NoError(t, err)

	assert.Equal(t, int32(4), resp2.Response.AffectedRows)

	resp3, err := service.FindRank(context.TODO(), &clubpb.FindRankRequest{
		Request: &typepb.Request{
			OrderBy: []*typepb.OrderBy{
				{Field: "name", Order: typepb.Order_DESC},
			},
		},
	})

	assert.NoError(t, err)

	assert.Equal(t, 4, len(resp3.Data))

	for i, d := range resp3.Data {
		if i == 2 {
			break
		}

		d.Name = d.Name + "1"
	}

	resp4, err := service.UpdateRank(context.TODO(), &clubpb.UpdateRankRequest{
		Data: resp3.Data,
	})

	assert.NoError(t, err)

	assert.Equal(t, int32(2), resp4.Response.AffectedRows)

	resp5, err := service.DeleteRank(context.TODO(), &clubpb.DeleteRankRequest{
		Id: array.Map[*clubpb.Rank, string](func(a *clubpb.Rank) string { return a.Id })(resp3.Data),
	})

	assert.NoError(t, err)

	assert.Equal(t, int32(4), resp5.Response.AffectedRows)
}
