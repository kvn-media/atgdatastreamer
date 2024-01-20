// internal/usecase/dataTank_usecase.go
package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/kvn-media/atgdatastreamer/internal/delivery"
	"github.com/kvn-media/atgdatastreamer/internal/models"
	"github.com/kvn-media/atgdatastreamer/internal/repository"
	"github.com/kvn-media/atgdatastreamer/internal/serial"
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

const serialReadTimeout = 5 * time.Second

// ReadFromSerial reads data from the serial port
func (u *dataTankUsecase) ReadFromSerial() ([]byte, error) {
	var receivedData []byte
	done := make(chan struct{})

	// Callback function to be executed when data is received
	callback := func(data []byte) {
		// Further processing of data (e.g., storing in the database)
		receivedData = append(receivedData, data...)
		// Signal that data has been received
		close(done)
	}

	// Start reading data with the provided callback
	go u.serialPort.StartReading(callback)

	// Wait for data to be received or an error
	select {
	case <-done:
		// Data received successfully
		// Further processing of receivedData (e.g., storing in the database)
		parsedData, parseErr := serial.ParseSS160PLUSProtocol(receivedData)
		if parseErr != nil {
			log.Printf("Error parsing received data: %v", parseErr)
			return nil, fmt.Errorf("failed to parse received data: %v", parseErr)
		}
		log.Println("Data received successfully from serial")
		return parsedData, nil

	case <-time.After(serialReadTimeout):
		// Timeout waiting for data
		log.Println("Timeout waiting for serial data in ReadFromSerial")
		return nil, fmt.Errorf("timeout waiting for serial data in ReadFromSerial")
	}
}

func (u *dataTankUsecase) WriteToSerial(data []byte) (int, error) {
	n, err := u.serialPort.Write(data)
	if err != nil {
		return 0, fmt.Errorf("failed to write to serial: %w", err)
	}
	return n, nil
}
