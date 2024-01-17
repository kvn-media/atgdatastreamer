package serial

import (
	"io"

	"github.com/tarm/serial"
)

type SerialPort interface {
	Connect(portName string, baudRate int) error
	Disconnect() error
	Read() ([]byte, error)
	Write(data []byte) (int, error)
}

// SerialPortImpl adalah implementasi dari SerialPort interface
type serialPort struct {
	port io.ReadWriteCloser
}

// NewSerialPortImpl inisialisasi SerialPortImpl
func NewSerialPortImpl() *serialPort {
	return &serialPort{}
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
	if s.port != nil {
		return s.port.Close()
	}
	return nil
}

// Read membaca data dari koneksi serial
func (s *serialPort) Read() ([]byte, error) {
	// Implementasi membaca data dari port serial
    buffer := make([]byte, 256)
    n, err := s.port.Read(buffer)
    if err!= nil {
        return nil, err
    }
    return buffer[:n], nil
}

// Write menulis data ke koneksi serial
func (s *serialPort) Write(data []byte) (int, error) {
	// Implementasi menulis data ke port serial
    n, err := s.port.Write(data)
    if err != nil {
        return 0, err
    }
    return n, nil
}
