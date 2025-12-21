package main

import (
	"log"
	"product/internal/grpc"
)

func main() {
	productService, err := grpc.InitGRPCServer()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC: %v", err)
	}

	log.Println("Starting gRPC server on :9090...")
	grpc.StartGRPCServer(productService)
}
