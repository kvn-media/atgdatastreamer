// internal/serial/ss160plus_mock.go

package serial

import "fmt"

// MockSerialPort is a mock implementation of the SerialPort interface
type MockSerialPort struct {
	DataToSend    []byte
	DataReceived  []byte
	ShouldSucceed bool
}

// Connect simulates the connection to a serial port
func (m *MockSerialPort) Connect(portName string, baudRate int) error {
	// Implement mock behavior if needed
	return nil
}

// Disconnect simulates disconnecting from a serial port
func (m *MockSerialPort) Disconnect() error {
	// Implement mock behavior if needed
	return nil
}

// StartReading simulates starting to read data from the serial port
func (m *MockSerialPort) StartReading(callback func([]byte)) {
	// Implement mock behavior if needed
}

// StopReading simulates stopping reading data from the serial port
func (m *MockSerialPort) StopReading() {
	// Implement mock behavior if needed
}

// Read simulates reading data from the serial port
func (m *MockSerialPort) Read(callback func([]byte)) {
	// Implement mock behavior
	if m.ShouldSucceed {
		callback(m.DataReceived)
	}
}

// Write simulates writing data to the serial port
func (m *MockSerialPort) Write(data []byte) (int, error) {
	// Implement mock behavior
	if m.ShouldSucceed {
		return len(data), nil
	}
	return 0, fmt.Errorf("mock Write error")
}

// ParseSS160PLUSProtocol mocks the SS160PLUS Protocol parsing logic
func (m *MockSerialPort) ParseSS160PLUSProtocol(data []byte) ([]byte, error) {
	// Implement mock behavior
	return data, nil
}
