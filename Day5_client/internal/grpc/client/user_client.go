package client

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo/pb/userpb"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (c *UserClient) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.User, error) {
	resp, err := c.client.CreateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}
