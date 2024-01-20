// internal/repository/dataTank_repository.go

package repository

import (
	"github.com/kvn-media/atgdatastreamer/internal/models"
	"gorm.io/gorm"
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
	db *gorm.DB
}

// NewDataTankRepository initializes DataTankRepository
func NewDataTankRepository(db *gorm.DB) (*dataTankRepo, error) {
	// Auto Migrate will create the table. If it already exists, it won't do anything.
	err := db.AutoMigrate(&models.DataTank{})
	if err != nil {
		return nil, err
	}

	return &dataTankRepo{
		db: db,
	}, nil
}

// CreateDataTank creates a new data tank in the database
func (r *dataTankRepo) CreateDataTank(dataTank *models.DataTank) error {
	return r.db.Create(dataTank).Error
}

// GetDataTanks retrieves all data tanks from the database
func (r *dataTankRepo) GetDataTanks() ([]*models.DataTank, error) {
	var dataTanks []*models.DataTank
	err := r.db.Find(&dataTanks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dataTanks, nil
}

// UpdateDataTank updates data in the database
func (r *dataTankRepo) UpdateDataTank(dataTank *models.DataTank) error {
	return r.db.Save(dataTank).Error
}

// DeleteDataTank deletes data from the database
func (r *dataTankRepo) DeleteDataTank(id int) error {
	return r.db.Delete(&models.DataTank{}, id).Error
}
