// internal/database/initDb.go

package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

// InitDB initializes the connection to SQLite and performs database migration
func InitDB(dbPath string) (*sql.DB, error) {
	var err error

	// Initialize the connection to SQLite
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
		return nil, err
	}

	log.Println("Connected to the database")

	// Perform database migration
	err = PerformDatabaseMigration(db, dbPath)
	if err != nil {
		log.Fatalf("Failed to perform database migration: %v", err)
		return nil, err
	}

	return db, nil
}

// CloseDB closes the database connection
func CloseDB(db *sql.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error while closing the database: %v", err)
		}
		log.Println("Connection to the database closed")
	}
}

// PerformDatabaseMigration performs database migration
func PerformDatabaseMigration(db *sql.DB, dbPath string) error {
	// Replace the migrationDir with the actual path in your project structure
	migrationDir := "internal/database/migration/000_init.sql"

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationDir,
		"sqlite3:///"+dbPath+"?_fk=true", // Updated SQLite driver URL
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}