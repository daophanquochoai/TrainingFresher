package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-demo/pb/productpb"
)

type ProductService struct {
	productpb.UnimplementedProductServiceServer
}

// urary
func (s *ProductService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.Product, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	return &productpb.Product{
		Id:    req.Id,
		Name:  "MacBook Pro",
		Price: "2500",
	}, nil
}

func (s *ProductService) ListProducts(
	ctx context.Context,
	_ *emptypb.Empty,
) (*productpb.ProductList, error) {

	return &productpb.ProductList{
		Product: []*productpb.Product{
			{Id: "1", Name: "Macbook", Price: "2000"},
			{Id: "2", Name: "iPad", Price: "900"},
		},
	}, nil
}
