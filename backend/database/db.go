package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getenv("DB_HOST", "localhost"),
		getenv("DB_PORT", "5432"),
		getenv("POSTGRES_USER", ""),
		getenv("POSTGRES_PASSWORD", ""),
		getenv("POSTGRES_DB", ""),
	)
}

func Connect() error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
