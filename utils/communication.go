package utils

import (
	"fmt"
	"net"
)

// CommunicationUtils provides utility functions for communication
type CommunicationUtils struct{}

// NewCommunicationUtils creates a new instance of CommunicationUtils
func NewCommunicationUtils() *CommunicationUtils {
	return &CommunicationUtils{}
}

// SendDataToMiniPC sends data to the mini PC over a network connection
func (cu *CommunicationUtils) SendDataToMiniPC(data []byte, miniPCAddress string) error {
	conn, err := net.Dial("tcp", miniPCAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to mini PC: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to send data to mini PC: %v", err)
	}

	return nil
}

// ReceiveDataFromConsole receives data from the console over a network connection
func (cu *CommunicationUtils) ReceiveDataFromConsole(consoleAddress string) ([]byte, error) {
	listen, err := net.Listen("tcp", consoleAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on console address: %v", err)
	}
	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		return nil, fmt.Errorf("failed to accept connection from console: %v", err)
	}
	defer conn.Close()

	// Read data from the console
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from console: %v", err)
	}

	return buffer[:n], nil
}
