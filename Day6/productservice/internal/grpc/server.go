package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"product/internal/grpc/interceptor"
	"product/internal/grpc/service"
	"product/pb/productservice/proto/productpb"
)

func StartGRPCServer(productService *service.ProductService) {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LoggingUnaryInterceptor,
		),
	)

	productpb.RegisterProductServiceServer(server, productService)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
