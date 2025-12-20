package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-demo/pb/userpb"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
}

// urary
func (s *UserService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.UserResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	user := &userpb.User{
		Id:    "1",
		Name:  req.Name + "grpc",
		Email: req.Email,
	}

	return &userpb.UserResponse{
		User: user,
	}, nil
}
