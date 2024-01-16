// internal/usecase/usecase.go

package usecase

import (
	"github.com/kvn-media/atgdatastreamer/internal/models"
	"github.com/kvn-media/atgdatastreamer/internal/repository"
	"github.com/kvn-media/atgdatastreamer/internal/serial"
)

// DataTankUsecase adalah usecase untuk entitas DataTank
type DataTankUsecase struct {
    dataTankRepository repository.DataTankRepository
	serialPort         serial.SerialPort
}

// NewDataTankUsecase inisialisasi DataTankUsecase
func NewDataTankUsecase(dataTankRepository repository.DataTankRepository, serialPort serial.SerialPort) *DataTankUsecase {
    return &DataTankUsecase{
		dataTankRepository: dataTankRepository,
		serialPort:         serialPort,
	}
}

// CreateDataTank membuat data baru
func (u *DataTankUsecase) CreateDataTank(dataTank *models.DataTank) error {
    return u.dataTankRepository.CreateDataTank(dataTank)
}

// GetDataTanks mengambil semua data
func (u *DataTankUsecase) GetDataTanks() ([]*models.DataTank, error) {
    return u.dataTankRepository.GetDataTanks()
}

// UpdateDataTank mengupdate data
func (u *DataTankUsecase) UpdateDataTank(dataTank *models.DataTank) error {
    return u.dataTankRepository.UpdateDataTank(dataTank)
}

// DeleteDataTank menghapus data
func (u *DataTankUsecase) DeleteDataTank(id int) error {
    return u.dataTankRepository.DeleteDataTank(id)
}

// ReadFromSerial membaca data dari port serial
func (u *DataTankUsecase) ReadFromSerial() ([]byte, error) {
    return u.serialPort.Read()
}

// WriteToSerial menulis data ke port serial
func (u *DataTankUsecase) WriteToSerial(data []byte) (int, error) {
    return u.serialPort.Write(data)
}