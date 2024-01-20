// internal/configs/config.go

package configs

import (
	"encoding/json"
	"os"
	"time"
)

// Config adalah struktur untuk menyimpan konfigurasi aplikasi
type Config struct {
	DBPath         string `json:"dbPath"`
	SerialPortName string `json:"serial_port_name"`
	SerialPortBaud int    `json:"serial_port_baud"`
	HTTPSEndpoint  string `json:"https_endpoint"`

	// New field for graceful shutdown timeout in seconds
	GracefulShutdownTimeout time.Duration `json:"graceful_shutdown_timeout"`
}

// LoadConfig membaca konfigurasi dari file eksternal
func LoadConfig(DBPath string) (Config, error) {
	var config Config

	file, err := os.Open(DBPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
