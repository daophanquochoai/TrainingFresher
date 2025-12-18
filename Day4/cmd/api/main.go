package main

import (
	"context"
	"flag"
	"fmt"
	"go-db-demo/internal/cache"
	"go-db-demo/internal/db"
	"go-db-demo/internal/handler"
	"go-db-demo/internal/middleware"
	"go-db-demo/internal/repository"
	"go-db-demo/internal/router"
	"go-db-demo/internal/service"
	"go-db-demo/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// declare env config
	configPath := flag.String("config", "../../config.yaml", "path to config file")
	flag.Parse()

	cfg, err := utils.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("App will start at port: %d", cfg.App.Port)
	log.Printf("Database URL: %s", cfg.DB.URL)

	// connect database
	database, err := db.Connect(cfg.DB.URL)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	log.Println("Successfully connected to the database")

	// redis
	redis, errRedis := cache.NewRedisClient(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.Db)
	if errRedis != nil {
		log.Fatal(errRedis)
	}
	defer redis.Close()
	// init
	userRepo := repository.NewPostgresUserRepository(database, redis)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHandler(userService)

	categoryRepo := repository.NewPostgresCategoryRepository(database, redis)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewPostgresProductRepository(database, redis)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	r := router.NewRouter(userHandler, categoryHandler, productHandler)

	// route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("UP"))
	})

	handlerMiddle := middleware.Chain(
		r,
		middleware.RateLimit(redis),
	)

	// create http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: handlerMiddle,
	}

	// run server in goroutine
	go func() {
		log.Printf("Server listening on port %d", cfg.App.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	if err := database.Close(); err != nil {
		log.Println("Error closing database:", err)
	}

	log.Println("Server exited gracefully")
}
