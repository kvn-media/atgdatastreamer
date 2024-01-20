// internal/database/initDb.go

package database

import (
	"log"

	"github.com/kvn-media/atgdatastreamer/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the connection to SQLite and performs database migration
func InitDB(dbPath string) (*gorm.DB, error) {
	var err error

	// Initialize the connection to SQLite using GORM
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
		return nil, err
	}

	// Perform automatic migration using GORM
	err = db.AutoMigrate(&models.DataTank{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
		return nil, err
	}

	// Ping the database to ensure the connection is established
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB: %v", err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
		return nil, err
	}

	log.Println("Connected to the database")

	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error while closing the database: %v", err)
		}

		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error while closing the database: %v", err)
		}
		log.Println("Connection to the database closed")
	}
}
