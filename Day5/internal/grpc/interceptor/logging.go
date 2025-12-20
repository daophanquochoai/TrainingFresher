package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func LoggingUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf(
		"[gRPC] method=%s duration=%s error=%v",
		info.FullMethod,
		time.Since(start),
		err,
	)

	return resp, err
}
