package main

import (
	"grpc-demo/internal/grpc/client"
	httpHandler "grpc-demo/internal/http"
	"log"
	"net/http"
)

func main() {
	conn := client.NewGRPCConn("localhost:9000")

	userClient := client.NewUserClient(conn)
	productClient := client.NewProductClient(conn)

	handler := httpHandler.NewHandler(userClient, productClient)

	http.HandleFunc("/users", handler.CreateUserHandler)
	http.HandleFunc("/products", handler.GetProductList)

	log.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", nil)
}
