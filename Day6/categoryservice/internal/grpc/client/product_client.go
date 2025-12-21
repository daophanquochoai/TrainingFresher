package client

import (
	"category/pb/productservice/proto/productpb"
	"context"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client productpb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: productpb.NewProductServiceClient(conn),
	}
}

func (c *ProductClient) GetProductByCategoryId(ctx context.Context, in string) (*productpb.ProductListResponse, error) {
	request := productpb.CategoryRequest{
		CategoryId: in,
	}
	resp, err := c.client.GetProductByCategoryId(ctx, &request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
