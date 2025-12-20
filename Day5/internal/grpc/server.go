package grpc

import (
	"google.golang.org/grpc"
	"grpc-demo/internal/grpc/interceptor"
	"grpc-demo/internal/grpc/service"
	"grpc-demo/pb/productpb"
	"grpc-demo/pb/userpb"

	"log"
	"net"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LoggingUnaryInterceptor,
			interceptor.AuthInterceptor,
		),
	)

	userpb.RegisterUserServiceServer(server, &service.UserService{})
	productpb.RegisterProductServiceServer(server, &service.ProductService{})

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
