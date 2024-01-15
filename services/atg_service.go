package services

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/kvn-media/atgdatastreamer/models"
)

func ReadATGData(conn net.Conn) ([]byte, error) {
	// Read data from ATG
	data := make([]byte, 1024)

	// Read data from ATG into the byte slice
	n, err := conn.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error reading from ATG: %w", err) // Wrap the error for context
	}

	// Return the actual data read
	return data[:n], nil // Return only the slice containing the read data
}

func ParseATGData(data []byte) (*models.ATGData, error) {
	// Parse data
	atgData := models.ATGData{}
	err := json.Unmarshal(data, &atgData)
	if err != nil {
		return nil, err
	}

	return &atgData, nil
}

func PrintATGData(data *models.ATGData) {
	fmt.Println("Type:", data.Type)
	fmt.Println("ID:", data.ID)
	fmt.Println("Value:", data.Value)
	fmt.Println("Timestamp:", data.Timestamp)
}
