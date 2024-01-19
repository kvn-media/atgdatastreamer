// internal/repository/dataTank_repository.go

package repository

import (
	"database/sql"

	"github.com/kvn-media/atgdatastreamer/internal/models"
)

// DataTankRepository adalah interface untuk entitas DataTank
type DataTankRepository interface {
	CreateDataTank(dataTank *models.DataTank) error
	GetDataTanks() ([]*models.DataTank, error)
	UpdateDataTank(dataTank *models.DataTank) error
	DeleteDataTank(id int) error
}

// DataTankRepository adalah repository untuk entitas DataTank
type dataTankRepo struct {
	db *sql.DB
}

// NewDataTankRepository inisialisasi DataTankRepository
func NewDataTankRepository(db *sql.DB) (*dataTankRepo, error) {
	// Check if the table exists, create it if not
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS data_tank (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    capacity INTEGER NOT NULL,
    temperature REAL
)
`)
	if err != nil {
		return nil, err
	}

	return &dataTankRepo{
		db: db,
	}, nil
}

// CreateDataTank membuat data baru di database
func (r *dataTankRepo) CreateDataTank(dataTank *models.DataTank) error {
	_, err := r.db.Exec("INSERT INTO data_tank (ID, Name, Capacity, Temperature) VALUES (?, ?, ?, ?)",
		dataTank.ID, dataTank.Name, dataTank.Capacity, dataTank.Temperature)
	return err
}

// GetDataTanks mengambil semua data dari database
func (r *dataTankRepo) GetDataTanks() ([]*models.DataTank, error) {
	rows, err := r.db.Query("SELECT ID, Name, Capacity, Temperature FROM data_tank")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataTanks []*models.DataTank
	for rows.Next() {
		var dt models.DataTank
		if err := rows.Scan(&dt.ID, &dt.Name, &dt.Capacity, &dt.Temperature); err != nil {
			return nil, err
		}
		dataTanks = append(dataTanks, &dt)
	}
	return dataTanks, nil
}

// UpdateDataTank mengupdate data di database
func (r *dataTankRepo) UpdateDataTank(dataTank *models.DataTank) error {
	_, err := r.db.Exec("UPDATE data_tank SET ID=?, Capacity=?, Temperature=? WHERE id=?",
		dataTank.ID, dataTank.Capacity, dataTank.Temperature, dataTank.ID)
	return err
}

// DeleteDataTank menghapus data dari database
func (r *dataTankRepo) DeleteDataTank(id int) error {
	_, err := r.db.Exec("DELETE FROM data_tank WHERE id=?", id)
	return err
}
