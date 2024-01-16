package configs

import (
	"log"
	"os"
	"strconv"
)

// Config holds application configurations
type Config struct {
	// Communication settings
	MiniPCAddress string // Alamat IP atau host Mini PC
	MiniPCPort    int    // Port untuk komunikasi dengan Mini PC

	// Database settings
	SQLiteDBPath string // Path untuk database SQLite

	// Add other configuration parameters as needed
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	miniPCAddress := os.Getenv("MINI_PC_ADDRESS")
	if miniPCAddress == "" {
		log.Fatal("MINI_PC_ADDRESS is not set")
	}

	miniPCPortStr := os.Getenv("MINI_PC_PORT")
	miniPCPort, err := strconv.Atoi(miniPCPortStr)
	if err != nil {
		log.Fatal("Invalid value for MINI_PC_PORT")
	}

	sqliteDBPath := os.Getenv("SQLITE_DB_PATH")
	if sqliteDBPath == "" {
		log.Fatal("SQLITE_DB_PATH is not set")
	}

	return &Config{
		MiniPCAddress: miniPCAddress,
		MiniPCPort:    miniPCPort,
		SQLiteDBPath:  sqliteDBPath,
		// Add other configuration parameters as needed
	}
}
