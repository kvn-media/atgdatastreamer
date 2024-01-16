package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kvn-media/atgdatastreamer/configs"
	"github.com/kvn-media/atgdatastreamer/managers"
	"github.com/kvn-media/atgdatastreamer/utils"
)

func main() {
	// Load configurations
	config := configs.LoadConfig()
	if config != nil {
		log.Fatalf("Error loading config: %v", config)
	}

	// Initialize communication utilities
	communicationUtils := utils.NewCommunicationUtils()

	// Initialize ATG Manager
	atgManager := managers.NewATGManager(communicationUtils)

	// Initialize signal handler for graceful shutdown
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-gracefulShutdown
		fmt.Printf("Received %s signal. Shutting down...\n", sig)
		// Perform cleanup or additional shutdown tasks if needed

		os.Exit(0)
	}()

	// Start the application based on its role
	switch config.AppRole {
	case "console":
		// Logic for the console application
		consoleApp := NewConsoleApp(atgManager, communicationUtils)
		consoleApp.Run()
	case "mini_pc":
		// Logic for the mini PC application
		miniPCApp := NewMiniPCApp(atgManager, communicationUtils)
		miniPCApp.Run()
	default:
		log.Fatalf("Invalid application role in config: %s", config.AppRole)
	}
}
