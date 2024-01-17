package database

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

// InitDB inisialisasi koneksi ke SQLite dan melakukan migrasi database
func InitDB(dbPath string) (*sql.DB, error) {
	var err error

	// Inisialisasi koneksi ke SQLite
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

	// Migrasi database
	err = PerformDatabaseMigration(db)
	if err != nil {
		log.Fatalf("Failed to perform database migration: %v", err)
		return nil, err
	}

	return db, nil
}

// CloseDB menutup koneksi database
func CloseDB(db *sql.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error while closing the database: %v", err)
		}
		log.Println("Connection to the database closed")
	}
}

// PerformDatabaseMigration melakukan migrasi database
func PerformDatabaseMigration(db *sql.DB) error {
	// Replace the migrationDir with the actual path in your project structure
	migrationDir := filepath.Join("path", "to", "migrations")

	driver, err := sqliteWithInstance(db)
	if err != nil {
		return err
	}

	srcDriver, err := file.WithInstance(&migrate.FileMigrationSource{
		Dir: migrationDir,
	}, &file.Driver{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", srcDriver, "sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// sqliteWithInstance mengembalikan instance driver SQLite
func sqliteWithInstance(db *sql.DB) (*sql.DB, error) {
	return db, nil
}
