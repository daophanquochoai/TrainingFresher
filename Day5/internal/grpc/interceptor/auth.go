package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func AuthInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	// check header
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("md: %+v\n", md)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}
	// get token
	authVals := md.Get("Authorization")
	if len(authVals) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization")
	}

	// parse token
	auth := authVals[0]
	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, status.Error(codes.Unauthenticated, "invalid authorization format")
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token != "123" {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	// token OK → đi tiếp vào handler
	return handler(ctx, req)
}
