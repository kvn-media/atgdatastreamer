package serial

import (
	"io"
	"log"

	"github.com/tarm/serial"
)

type SerialPort interface {
	Connect(portName string, baudRate int) error
	Disconnect() error
	StartReading(callback func([]byte))
	StopReading()
	Read(callback func([]byte))
	Write(data []byte) (int, error)
}

// SerialPortImpl adalah implementasi dari SerialPort interface
type serialPort struct {
	port     io.ReadWriteCloser
	readChan chan []byte
	stopChan chan struct{}
}

// NewSerialPortImpl inisialisasi SerialPortImpl
func NewSerialPortImpl() *serialPort {
	return &serialPort{
		readChan: make(chan []byte),
		stopChan: make(chan struct{}),
	}
}

// Connect membuka koneksi serial
func (s *serialPort) Connect(portName string, baudRate int) error {
	config := &serial.Config{Name: portName, Baud: baudRate}
	port, err := serial.OpenPort(config)
	if err != nil {
		return err
	}
	s.port = port
	return nil
}

// Disconnect menutup koneksi serial
func (s *serialPort) Disconnect() error {
	close(s.stopChan) // Signal to stop reading goroutine
	if s.port != nil {
		err := s.port.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *serialPort) StartReading(callback func([]byte)) {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return // Stop goroutine when signaled
			default:
				buffer := make([]byte, 256)
				n, err := s.port.Read(buffer)
				if err != nil {
					log.Printf("Error reading from serial port: %v", err)
					return
				}
				callback(buffer[:n])
			}
		}
	}()
}

func (s *serialPort) StopReading() {
	close(s.stopChan) // Signal to stop reading goroutine
}

// Read membaca data dari koneksi serial
func (s *serialPort) Read(callback func([]byte)) {
	buffer := make([]byte, 256)
	n, err := s.port.Read(buffer)
	if err != nil {
		log.Printf("Error reading from serial port: %v", err)
		return
	}
	callback(buffer[:n])
}

// Write menulis data ke koneksi serial
func (s *serialPort) Write(data []byte) (int, error) {
	n, err := s.port.Write(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}
