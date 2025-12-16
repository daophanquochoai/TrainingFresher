package main

import (
	"fmt"
	"go-rest-api/internal/handler"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/service"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	// health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("UP"))
	})

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/", userHandler.GetUserByID)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
