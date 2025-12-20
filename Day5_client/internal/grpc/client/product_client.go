package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-demo/pb/productpb"
)

type ProductClient struct {
	client productpb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: productpb.NewProductServiceClient(conn),
	}
}

func (c *ProductClient) GetProduct(ctx context.Context, id string) (*productpb.Product, error) {
	return c.client.GetProduct(ctx, &productpb.GetProductRequest{Id: id})
}

func (c *ProductClient) ListProducts(ctx context.Context) ([]*productpb.Product, error) {
	resp, err := c.client.ListProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Product, nil
}
