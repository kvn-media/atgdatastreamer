package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// ATGRepository handles interactions with the ATG data in the SQLite database
type ATGRepository struct {
	DB *sql.DB
}

// NewATGRepository creates a new instance of ATGRepository
func NewATGRepository(dbPath string) (*ATGRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Set up the database schema if needed
	if err := setupDatabaseSchema(db); err != nil {
		return nil, err
	}

	return &ATGRepository{
		DB: db,
	}, nil
}

// setupDatabaseSchema creates the necessary tables if they don't exist
func setupDatabaseSchema(db *sql.DB) error {
	// Example: Creating ATGData table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS ATGData (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			// Add other fields as needed
		)
	`)
	return err
}

// SaveATGData saves ATG data to the database
func (r *ATGRepository) SaveATGData(data interface{}) error {
	// Implement logic to save ATG data to the database
	// Example: Insert data into the ATGData table
	query := "INSERT INTO ATGData (Timestamp) VALUES (?)"
	_, err := r.DB.Exec(query, time.Now())
	return err
}

// GetLatestATGData retrieves the latest ATG data from the database
func (r *ATGRepository) GetLatestATGData() (interface{}, error) {
	// Implement logic to retrieve the latest ATG data from the database
	// Example: Select the most recent entry from the ATGData table
	query := "SELECT * FROM ATGData ORDER BY Timestamp DESC LIMIT 1"
	row := r.DB.QueryRow(query)

	var result interface{} // Change this to match your data structure
	err := row.Scan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no ATG data found")
		}
		return nil, err
	}

	return result, nil
}
