package database

import (
	"log"

	"github.com/phzeng0726/go-graphql-server-template/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// TODO: You can change to another SQL driver if needed.
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("Database connected")
	return conn
}

// Ensure that the database and model formats match
func SyncDatabase(conn *gorm.DB) {
	err := conn.AutoMigrate(&domain.Group{})
	if err != nil {
		log.Fatalf("Failed to migrate Group: %v", err)
	}

	err = conn.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User: %v", err)
	}
}
