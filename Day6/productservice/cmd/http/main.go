package main

import (
	"flag"
	"fmt"
	"log"
	"product/internal"
)

func main() {
	port := flag.Int("port", 8080, "Server port")
	flag.Parse()

	app, err := internal.New()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server starting on %s", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
