package v1

import (
	"context"
	"estore/protoc/stockpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StockService struct {
	stockpb.UnsafeStockSystemServer
}

func (StockService) Ping(ctx context.Context, req *stockpb.PingPong) (*stockpb.PingPong, error) {
	req.Message = "Pong"
	return req, nil
}

func (StockService) CreateCommodity(ctx context.Context, req *stockpb.CreateCommodityRequest) (*stockpb.CreateCommodityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommodity not implemented")
}
func (StockService) UpdateCommodity(ctx context.Context, req *stockpb.UpdateCommodityRequest) (*stockpb.UpdateCommodityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommodity not implemented")
}
func (StockService) FindCommodity(ctx context.Context, req *stockpb.FindCommodityRequest) (*stockpb.FindCommodityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCommodity not implemented")
}
func (StockService) CreateSpec(ctx context.Context, req *stockpb.CreateSpecRequest) (*stockpb.CreateSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpec not implemented")
}
func (StockService) UpdateSpec(ctx context.Context, req *stockpb.UpdateSpecRequest) (*stockpb.UpdateSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpec not implemented")
}
func (StockService) FindSpec(ctx context.Context, req *stockpb.FindSpecRequest) (*stockpb.FindSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSpec not implemented")
}
func (StockService) CreateSupplier(ctx context.Context, req *stockpb.CreateSupplierRequest) (*stockpb.CreateSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSupplier not implemented")
}
func (StockService) UpdateSupplier(ctx context.Context, req *stockpb.UpdateSupplierRequest) (*stockpb.UpdateSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSupplier not implemented")
}
func (StockService) DeleteSupplier(ctx context.Context, req *stockpb.DeleteSupplierRequest) (*stockpb.DeleteSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSupplier not implemented")
}
func (StockService) FindSupplier(ctx context.Context, req *stockpb.FindSupplierRequest) (*stockpb.FindSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSupplier not implemented")
}
func (StockService) CreateStorehouse(ctx context.Context, req *stockpb.CreateStorehouseRequest) (*stockpb.CreateStorehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorehouse not implemented")
}
func (StockService) UpdateStorehouse(ctx context.Context, req *stockpb.UpdateStorehouseRequest) (*stockpb.UpdateStorehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorehouse not implemented")
}
func (StockService) FindStorehouse(ctx context.Context, req *stockpb.FindStorehouseRequest) (*stockpb.FindStorehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStorehouse not implemented")
}
func (StockService) CreatePurchase(ctx context.Context, req *stockpb.CreatePurchaseRequest) (*stockpb.CreatePurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePurchase not implemented")
}
func (StockService) UpdatePurchase(ctx context.Context, req *stockpb.UpdatePurchaseRequest) (*stockpb.UpdatePurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePurchase not implemented")
}
func (StockService) FindPurchase(ctx context.Context, req *stockpb.FindPurchaseRequest) (*stockpb.FindPurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPurchase not implemented")
}
func (StockService) UpdateInventory(ctx context.Context, req *stockpb.UpdateInventoryRequest) (*stockpb.UpdateInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInventory not implemented")
}
func (StockService) FindInventory(ctx context.Context, req *stockpb.FindInventoryRequest) (*stockpb.FindInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindInventory not implemented")
}
