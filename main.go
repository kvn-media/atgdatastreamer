package main

import (
	"log"

	"github.com/kvn-media/atgdatastreamer/internal/application"
	"github.com/kvn-media/atgdatastreamer/internal/configs"
	"github.com/kvn-media/atgdatastreamer/internal/controllers"
	"github.com/kvn-media/atgdatastreamer/internal/database"
	"github.com/kvn-media/atgdatastreamer/internal/delivery"
	"github.com/kvn-media/atgdatastreamer/internal/repository"
	"github.com/kvn-media/atgdatastreamer/internal/serial"
	"github.com/kvn-media/atgdatastreamer/internal/usecase"
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig("internal/configs/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database
	db, err := database.InitDB(config.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.CloseDB(db)

	// Create repository
	dataTankRepository, err := repository.NewDataTankRepository(db)
	if err != nil {
		log.Fatalf("Failed to create data tank repository: %v", err)
	}

	// Create serial port
	serialPort := serial.NewSerialPortImpl()
	err = serialPort.Connect(config.SerialPortName, config.SerialPortBaud)
	if err != nil {
		log.Fatalf("Failed to connect serial port: %v", err)
	}
	defer serialPort.Disconnect()

	// Create HTTPS delivery
	httpsDelivery := delivery.NewHttpsDelivery(config.HTTPSEndpoint)

	// Create use case
	dataTankUsecase := usecase.NewDataTankUsecase(dataTankRepository, serialPort, httpsDelivery)

	// Create controller
	dataTankController := controllers.NewDataTankController(dataTankUsecase)

	// Create and run the application
	app := application.NewApp(dataTankController, config)
	app.Initialize()
	app.Run()
}