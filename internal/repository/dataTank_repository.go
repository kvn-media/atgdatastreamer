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
func NewDataTankRepository(db *sql.DB) *dataTankRepo {
	return &dataTankRepo{
		db: db,
	}
}

// CreateDataTank membuat data baru di database
func (r *dataTankRepo) CreateDataTank(dataTank *models.DataTank) error {
	_, err := r.db.Exec("INSERT INTO data_tank (field1, field2, field3) VALUES (?, ?, ?)",
		dataTank.Field1, dataTank.Field2, dataTank.Field3)
	return err
}

// GetDataTanks mengambil semua data dari database
func (r *dataTankRepo) GetDataTanks() ([]*models.DataTank, error) {
	rows, err := r.db.Query("SELECT field1, field2, field3 FROM data_tank")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataTanks []*models.DataTank
	for rows.Next() {
		var dt models.DataTank
		if err := rows.Scan(&dt.Field1, &dt.Field2, &dt.Field3); err != nil {
			return nil, err
		}
		dataTanks = append(dataTanks, &dt)
	}
	return dataTanks, nil
}

// UpdateDataTank mengupdate data di database
func (r *dataTankRepo) UpdateDataTank(dataTank *models.DataTank) error {
	_, err := r.db.Exec("UPDATE data_tank SET field1=?, field2=?, field3=? WHERE id=?",
		dataTank.Field1, dataTank.Field2, dataTank.Field3, dataTank.ID)
	return err
}

// DeleteDataTank menghapus data dari database
func (r *dataTankRepo) DeleteDataTank(id int) error {
	_, err := r.db.Exec("DELETE FROM data_tank WHERE id=?", id)
	return err
}
