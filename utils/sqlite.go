package utils

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitializeSQLite initializes SQLite database
func InitializeSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("your-database-file.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to SQLite database:", err)
		return nil, err
	}

	// Auto Migrate the models
	// Add your models here using db.AutoMigrate(&YourModel{})
	db.AutoMigrate()

	return db, nil
}
