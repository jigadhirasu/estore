package v1

import (
	"context"
	"estore/protoc/clubpb"

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

func (ClubService) CreateClub(ctx context.Context, req *clubpb.CreateClubRequest) (*clubpb.CreateClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClub not implemented")
}
func (ClubService) UpdateClub(ctx context.Context, req *clubpb.UpdateClubRequest) (*clubpb.UpdateClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClub not implemented")
}
func (ClubService) FindCluib(ctx context.Context, req *clubpb.FindClubRequest) (*clubpb.FindClubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCluib not implemented")
}
func (ClubService) CreateRank(ctx context.Context, req *clubpb.CreateRankRequest) (*clubpb.CreateRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRank not implemented")
}
func (ClubService) UpdateRank(ctx context.Context, req *clubpb.UpdateRankRequest) (*clubpb.UpdateRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRank not implemented")
}
func (ClubService) DeleteRank(ctx context.Context, req *clubpb.DeleteRankRequest) (*clubpb.DeleteRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRank not implemented")
}
func (ClubService) FindRank(ctx context.Context, req *clubpb.FindRankRequest) (*clubpb.FindRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRank not implemented")
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
