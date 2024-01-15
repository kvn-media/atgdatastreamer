package repositories

import (
	"database/sql"
	"github.com/kvn-media/atgdatastreamer/models"
)

type ATGRepository struct {
	db *sql.DB
}

func NewATGRepository(db *sql.DB) *ATGRepository {
	return &ATGRepository{db}
}

func (repository *ATGRepository) SaveATGData(data *models.ATGData) error {
	// Insert data into database
	stmt, err := repository.db.Prepare("INSERT INTO atg_data (type, tank_id, level, temperature, timestamp) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.Type, data.ID, data.Value, data.Timestamp)
	if err != nil {
		return err
	}

	return nil
}
