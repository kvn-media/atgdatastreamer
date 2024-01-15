package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ATGAddress string
	ATGPort    int
	ATGFormat  string
}

func LoadConfig() *Config {
	// Load configuration file
	config := &Config{}
	fileData, err := os.ReadFile("config.json") // Read the file and handle errors
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil
	}

	err = json.Unmarshal(fileData, config) // Use the file data for unmarshalling
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	// Check configuration
	if config.ATGAddress == "" {
		fmt.Println("ATG address must be specified")
		return nil
	}

	if config.ATGPort < 1 {
		fmt.Println("ATG port must be greater than or equal to 1")
		return nil
	}

	if config.ATGFormat != "json" {
		fmt.Println("ATG format must be JSON")
		return nil
	}

	return config
}
