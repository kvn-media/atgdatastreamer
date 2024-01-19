// internal/configs/config.go

package configs

import (
	"encoding/json"
	"os"
)

// Config adalah struktur untuk menyimpan konfigurasi aplikasi
type Config struct {
	DBPath         string `json:"db_path"`
	SerialPortName string `json:"serial_port_name"`
	SerialPortBaud int    `json:"serial_port_baud"`
	HTTPSEndpoint  string `json:"https_endpoint"`
}

// LoadConfig membaca konfigurasi dari file eksternal
func LoadConfig(path string) (Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
