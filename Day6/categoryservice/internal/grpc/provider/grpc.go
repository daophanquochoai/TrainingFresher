package provider

import (
	"category/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func ProvideGRPCConnection(cfg *config.Config) (*grpc.ClientConn, func(), error) {
	// Lấy địa chỉ từ config
	productServiceAddr := cfg.App.Grpc.Host + ":" + cfg.App.Grpc.Port
	if productServiceAddr == "" {
		productServiceAddr = "localhost:50051"
	}

	log.Printf("Connecting to Product Service at: %s", productServiceAddr)

	conn, err := grpc.Dial(
		productServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("Failed to connect to Product Service: %v", err)
		return nil, nil, err
	}

	log.Printf("Successfully connected to Product Service")

	// Cleanup function
	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close gRPC connection: %v", err)
		} else {
			log.Printf("gRPC connection closed")
		}
	}

	return conn, cleanup, nil
}
