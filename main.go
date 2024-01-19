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
	// Create instances of dependencies
	config, err := configs.LoadConfig("internal/configs/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.InitDB(config.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	dataTankRepository := repository.NewDataTankRepository(db)

	serialPort := serial.NewSerialPortImpl()
	err = serialPort.Connect(config.SerialPortName, config.SerialPortBaud)
	if err != nil {
		log.Fatalf("Failed to connect serial port: %v", err)
	}

	httpsDelivery := delivery.NewHttpsDelivery(config.HTTPSEndpoint)

	// Create an instance of DataTankUsecase and pass dependencies to it
	dataTankUsecase := usecase.NewDataTankUsecase(dataTankRepository, serialPort, httpsDelivery)

	// Create an instance of DataTankController and pass DataTankUsecase to it
	dataTankController := controllers.NewDataTankController(dataTankUsecase)

	// Create an instance of App and pass DataTankController to it
	app := application.NewApp(dataTankController, config)

	// Initialize and run the application
	app.Initialize()
	app.Run()
}