// internal/usecase/usecase.go

package usecase

import (
	"github.com/kvn-media/atgdatastreamer/internal/models"
	"github.com/kvn-media/atgdatastreamer/internal/repository"
	"github.com/kvn-media/atgdatastreamer/internal/serial"
	"github.com/kvn-media/atgdatastreamer/internal/delivery"
)

type DataTankUsecase interface {
	CreateDataTank(dataTank *models.DataTank) error
	GetDataTanks() ([]*models.DataTank, error)
	UpdateDataTank(dataTank *models.DataTank) error
	DeleteDataTank(id int) error
	ReadFromSerial() ([]byte, error)
	WriteToSerial(data []byte) (int, error) 
}

// DataTankUsecase adalah usecase untuk entitas DataTank
type dataTankUsecase struct {
    dataTankRepository repository.DataTankRepository
	serialPort         serial.SerialPort
	httpsDelivery      delivery.Delivery
}

// NewDataTankUsecase inisialisasi DataTankUsecase
func NewDataTankUsecase(dataTankRepository repository.DataTankRepository, serialPort serial.SerialPort, httpsDelivery delivery.Delivery) *dataTankUsecase {
	return &dataTankUsecase{
		dataTankRepository: dataTankRepository,
		serialPort:         serialPort,
		httpsDelivery:      httpsDelivery,
	}
}

// CreateDataTank membuat data baru
func (u *dataTankUsecase) CreateDataTank(dataTank *models.DataTank) error {
    return u.dataTankRepository.CreateDataTank(dataTank)
}

// GetDataTanks mengambil semua data
func (u *dataTankUsecase) GetDataTanks() ([]*models.DataTank, error) {
    return u.dataTankRepository.GetDataTanks()
}

// UpdateDataTank mengupdate data
func (u *dataTankUsecase) UpdateDataTank(dataTank *models.DataTank) error {
    return u.dataTankRepository.UpdateDataTank(dataTank)
}

// DeleteDataTank menghapus data
func (u *dataTankUsecase) DeleteDataTank(id int) error {
    return u.dataTankRepository.DeleteDataTank(id)
}

// ReadFromSerial membaca data dari port serial
func (u *dataTankUsecase) ReadFromSerial() ([]byte, error) {
    return u.serialPort.Read()
}

// WriteToSerial menulis data ke port serial
func (u *dataTankUsecase) WriteToSerial(data []byte) (int, error) {
    return u.serialPort.Write(data)
}