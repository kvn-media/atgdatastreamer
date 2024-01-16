package serial

import "io"

// SerialPort adalah interface untuk koneksi serial
type SerialPort interface {
    Connect(portName string, baudRate int) error
    Disconnect() error
    Read() ([]byte, error)
    Write(data []byte) (int, error)
}