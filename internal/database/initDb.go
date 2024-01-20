// internal/database/initDb.go

package database

import (
	"log"

	"github.com/kvn-media/atgdatastreamer/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// RunMigrations runs all database migrations
func RunMigrations(db *gorm.DB) {
	// AutoMigrate will create the table if it does not exist and add missing fields
	db.AutoMigrate(&models.DataTank{})

	// // You can also use `Migrator` to perform more complex migrations
	// migrator := db.Migrator()
	// if migrator.HasTable(&models.DataTank{}) {
	// 	// Add your update statements here if needed
	// 	// For example, if you want to add a new column named 'NewColumn'
	// 	// you can use the following:
	// 	// migrator.AddColumn(&models.DataTank{}, "NewColumn", &models.DataTank{}.NewColumn)
	// }

	// Add more AutoMigrate or update statements for other models if needed
}

// TruncateTables truncates the specified tables
func TruncateTables(db *gorm.DB, tables ...interface{}) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, table := range tables {
			if err := tx.Exec("DELETE FROM ?", table).Error; err != nil {
				return err
			}
			// If you're using SQLite, you might need to reset the auto-increment counter
			// Uncomment the line below if needed
			if err := tx.Exec("DELETE FROM sqlite_sequence WHERE name=?", table).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error truncating tables: %v", err)
		return err
	}
	log.Println("Tables truncated successfully")
	return nil
}

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
	RunMigrations(db)

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
