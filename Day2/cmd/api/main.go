package main

import (
	"flag"
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/handler"
	"go-db-demo/internal/repository"
	"go-db-demo/internal/router"
	"go-db-demo/internal/service"
	"go-db-demo/internal/utils"
	"log"
	"net/http"
)

func main() {
	configPath := flag.String("config", "../../config.yaml", "path to config file")
	flag.Parse()

	cfg, err := utils.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("App will start at port: %d", cfg.App.Port)
	log.Printf("Database URL: %s", cfg.DB.URL)
	database, err := db.Connect(cfg.DB.URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database")

	// init
	userRepo := repository.NewPostgresUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHandler(userService)

	categoryRepo := repository.NewPostgresCategoryRepository(database)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewPostgresProductRepository(database)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	r := router.NewRouter(userHandler, categoryHandler, productHandler)

	// route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("UP"))
	})

	//http.HandleFunc("POST /users", handler.CreateUser)
	//http.HandleFunc("GET /users/", handler.GetUserByID)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), r); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
