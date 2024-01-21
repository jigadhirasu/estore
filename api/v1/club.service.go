package v1

import (
	"context"
	"estore/internal/club"
	"estore/internal/mariadb"
	"estore/internal/rank"
	"estore/protoc/clubpb"

	"github.com/jigadhirasu/rex"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClubService struct {
	clubpb.UnsafeClubSystemServer
}

func (ClubService) Ping(ctx context.Context, req *clubpb.PingPong) (*clubpb.PingPong, error) {
	req.Message = "Pong"
	return req, nil
}

func (ClubService) CreateClub(cin context.Context, req *clubpb.CreateClubRequest) (*clubpb.CreateClubResponse, error) {
	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*clubpb.Club](dbname)),
		rex.Once(mariadb.F0Begin[*clubpb.Club](dbname)),
		club.PipeCreateClub(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*clubpb.CreateClubResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()

}

func (ClubService) UpdateClub(cin context.Context, req *clubpb.UpdateClubRequest) (*clubpb.UpdateClubResponse, error) {
	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*clubpb.Club](dbname)),
		rex.Once(mariadb.F0Begin[*clubpb.Club](dbname)),
		club.PipeUpdateClub(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*clubpb.UpdateClubResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()

}

func (ClubService) FindClub(cin context.Context, req *clubpb.FindClubRequest) (*clubpb.FindClubResponse, error) {

	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe2(
		rex.Once(mariadb.F0GetDB[*clubpb.FindClubRequest](dbname)),
		club.PipeFindClub(mariadb.DBKey(dbname)),
	)(rex.From(req))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()

}

func (ClubService) CreateRank(cin context.Context, req *clubpb.CreateRankRequest) (*clubpb.CreateRankResponse, error) {
	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*clubpb.Rank](dbname)),
		rex.Once(mariadb.F0Begin[*clubpb.Rank](dbname)),
		rank.PipeCreateRank(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*clubpb.CreateRankResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()

}

func (ClubService) UpdateRank(cin context.Context, req *clubpb.UpdateRankRequest) (*clubpb.UpdateRankResponse, error) {
	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[*clubpb.Rank](dbname)),
		rex.Once(mariadb.F0Begin[*clubpb.Rank](dbname)),
		rank.PipeUpdateRank(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*clubpb.UpdateRankResponse](dbname)),
	)(rex.From(req.Data...))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()
}

func (ClubService) DeleteRank(cin context.Context, req *clubpb.DeleteRankRequest) (*clubpb.DeleteRankResponse, error) {
	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe4(
		rex.Once(mariadb.F0GetDB[string](dbname)),
		rex.Once(mariadb.F0Begin[string](dbname)),
		rank.PipeDeleteRank(mariadb.TxKey(dbname)),
		rex.Tap(mariadb.F0Commit[*clubpb.DeleteRankResponse](dbname)),
	)(rex.From(req.Id...))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()
}

func (ClubService) FindRank(cin context.Context, req *clubpb.FindRankRequest) (*clubpb.FindRankResponse, error) {
	dbname := "test"

	ctx := rex.NewContext(cin)

	pipe := rex.Pipe2(
		rex.Once(mariadb.F0GetDB[*clubpb.FindRankRequest](dbname)),
		rank.PipeFindClub(mariadb.DBKey(dbname)),
	)(rex.From(req))(ctx)

	defer mariadb.Rollback(ctx, dbname)

	item := <-pipe()

	return item()
}

func (ClubService) CreateCustomer(ctx context.Context, req *clubpb.CreateCustomerRequest) (*clubpb.CreateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}

func (ClubService) UpdateCustomer(ctx context.Context, req *clubpb.UpdateCustomerRequest) (*clubpb.UpdateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}

func (ClubService) FindCustomer(ctx context.Context, req *clubpb.FindCustomerRequest) (*clubpb.FindCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCustomer not implemented")
}
