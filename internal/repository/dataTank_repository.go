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
    time timestamp,
    Barel BIGINT NOT NULL,
	VolumeBarel INTEGER NOT NULL,
    AveTemperature INTEGER NOT NULL,
    WaterDebit DECIMAL(10, 2) NOT NULL,
	TempProduct INTEGER NOT NULL,
    Alarm varchar(5000) NOT NULL
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
	_, err := r.db.Exec("INSERT INTO data_tank (ID, Time, Barel, VolumeBarel, AveTemperature, WaterDebit, TempProduct, Alarm) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		dataTank.ID, dataTank.Time, dataTank.Barel, dataTank.VolumeBarel, dataTank.AveTemperature, dataTank.WaterDebit, dataTank.TempProduct, dataTank.Alarm)
	return err
}

// GetDataTanks mengambil semua data dari database
func (r *dataTankRepo) GetDataTanks() ([]*models.DataTank, error) {
	rows, err := r.db.Query("SELECT ID, Time, Barel, VolumeBarel, AveTemperature, WaterDebit, TempProduct, Alarm FROM data_tank")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataTanks []*models.DataTank
	for rows.Next() {
		var dt models.DataTank
		if err := rows.Scan(&dt.ID, &dt.Time, &dt.Barel, &dt.VolumeBarel, &dt.AveTemperature, &dt.WaterDebit, &dt.TempProduct, &dt.Alarm); err != nil {
			return nil, err
		}
		dataTanks = append(dataTanks, &dt)
	}
	return dataTanks, nil
}

// UpdateDataTank mengupdate data di database
func (r *dataTankRepo) UpdateDataTank(dataTank *models.DataTank) error {
	_, err := r.db.Exec("UPDATE data_tank SET ID=?, Barel=?, VolumeBarel=?, AveTemperature=?, WaterDebit=?, TempProduct=?, Alarm=? WHERE ID=?",
		dataTank.ID, dataTank.Barel, dataTank.VolumeBarel, dataTank.AveTemperature, dataTank.WaterDebit, dataTank.TempProduct, dataTank.Alarm, dataTank.ID)
	return err
}

// DeleteDataTank menghapus data dari database
func (r *dataTankRepo) DeleteDataTank(id int) error {
	_, err := r.db.Exec("DELETE FROM data_tank WHERE id=?", id)
	return err
}
