// internal/usecase/dataTank_usecase.go
package usecase

import (
	"fmt"

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

type dataTankUsecase struct {
    dataTankRepository repository.DataTankRepository
	serialPort         serial.SerialPort
	httpsDelivery      delivery.Delivery
}

func NewDataTankUsecase(dataTankRepository repository.DataTankRepository, serialPort serial.SerialPort, httpsDelivery delivery.Delivery) *dataTankUsecase {
	return &dataTankUsecase{
		dataTankRepository: dataTankRepository,
		serialPort:         serialPort,
		httpsDelivery:      httpsDelivery,
	}
}

func (u *dataTankUsecase) CreateDataTank(dataTank *models.DataTank) error {
    if err := u.dataTankRepository.CreateDataTank(dataTank); err != nil {
        return fmt.Errorf("failed to create DataTank: %w", err)
    }
    return nil
}

func (u *dataTankUsecase) GetDataTanks() ([]*models.DataTank, error) {
    dataTanks, err := u.dataTankRepository.GetDataTanks()
    if err != nil {
        return nil, fmt.Errorf("failed to fetch DataTanks: %w", err)
    }
    return dataTanks, nil
}

func (u *dataTankUsecase) UpdateDataTank(dataTank *models.DataTank) error {
    if err := u.dataTankRepository.UpdateDataTank(dataTank); err != nil {
        return fmt.Errorf("failed to update DataTank: %w", err)
    }
    return nil
}

func (u *dataTankUsecase) DeleteDataTank(id int) error {
    if err := u.dataTankRepository.DeleteDataTank(id); err != nil {
        return fmt.Errorf("failed to delete DataTank: %w", err)
    }
    return nil
}

func (u *dataTankUsecase) ReadFromSerial() ([]byte, error) {
    data, err := u.serialPort.Read()
    if err != nil {
        return nil, fmt.Errorf("failed to read from serial: %w", err)
    }
    return data, nil
}

func (u *dataTankUsecase) WriteToSerial(data []byte) (int, error) {
    n, err := u.serialPort.Write(data)
    if err != nil {
        return 0, fmt.Errorf("failed to write to serial: %w", err)
    }
    return n, nil
}