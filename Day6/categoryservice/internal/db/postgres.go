package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg map[string]any) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg["host"], cfg["user"], cfg["password"],
		cfg["dbname"], cfg["port"], cfg["sslmode"],
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
