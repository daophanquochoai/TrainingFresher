package main

import (
	"flag"
	"fmt"
	"log"
	"product/internal/config"
)

func main() {
	// declare config
	configPath := flag.String("config", "./config/config.yaml", "Path to config file")
	flag.Parse()

	var cfg, err = config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("config: %+v\n", cfg)

	//// connect db
	//postgresConfig := db.PostgresConfig{
	//	Host: cfg.DB.Host,
	//	Port: cfg.DB.Port,
	//	User: cfg.DB.User,
	//	Pass: cfg.DB.Pass,
	//	Db:   cfg.DB.Name,
	//}
	//database, err := db.Connect(postgresConfig)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
}
