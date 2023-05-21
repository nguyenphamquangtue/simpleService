package pg

import (
	"errors"
	"fmt"
	"os"
	"simpleService/external/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Connection establishes a new database connection
func Connect() error {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	port := 5432
	sslMode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, username, password, database, port, sslMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect database %v", err))
	}

	// Migrate
	err = db.AutoMigrate(&dto.User{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to migrate database %v", err))
	}

	return nil
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
