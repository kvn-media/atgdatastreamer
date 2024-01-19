// internal/application/app.go
package application

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kvn-media/atgdatastreamer/internal/configs"
	"github.com/kvn-media/atgdatastreamer/internal/controllers"
	"github.com/kvn-media/atgdatastreamer/internal/database"
	"github.com/kvn-media/atgdatastreamer/internal/delivery"
	"github.com/kvn-media/atgdatastreamer/internal/repository"
	"github.com/kvn-media/atgdatastreamer/internal/serial"
	"github.com/kvn-media/atgdatastreamer/internal/usecase"

)

type App struct {
	router             *mux.Router
	server             *http.Server
	db                 *sql.DB
	dataTankController controllers.DataTankController
	config             configs.Config
}

func NewApp(dataTankController controllers.DataTankController, config configs.Config) *App {
	return &App{
		dataTankController: dataTankController,
		config:             config,
	}
}

func (app *App) Initialize() {
	// Initialize the database
	var err error
	app.db, err = database.InitDB(app.config.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Defer closing the database connection at the end of the application lifecycle
	defer database.CloseDB(app.db)

	// Migrate Database
	log.Println("Performing database migration...")
	err = database.PerformDatabaseMigration(app.db, app.config.DBPath)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration successful")

	// Initialize repository and usecase
	dataTankRepository, err := repository.NewDataTankRepository(app.db)
	if err != nil {
		log.Fatalf("Failed to create data tank repository: %v", err)
	}
	serialPort := serial.NewSerialPortImpl()

	err = serialPort.Connect(app.config.SerialPortName, app.config.SerialPortBaud)
	if err != nil {
		log.Fatalf("Failed to connect to the serial port: %v", err)
	}
	defer serialPort.Disconnect()

	// Initialize HTTPS Delivery
	httpsDelivery := delivery.NewHttpsDelivery(app.config.HTTPSEndpoint)

	dataTankUsecase := usecase.NewDataTankUsecase(dataTankRepository, serialPort, httpsDelivery)
	app.dataTankController = controllers.NewDataTankController(dataTankUsecase)

	// Initialize the router
	app.router = mux.NewRouter()
	app.initializeRoutes()
}

// initializeRoutes adds routes to the router
func (app *App) initializeRoutes() {
	app.router.HandleFunc("/data-tank", app.dataTankController.CreateDataTank).Methods("POST")
	app.router.HandleFunc("/data-tank", app.dataTankController.GetDataTanks).Methods("GET")
	app.router.HandleFunc("/data-tank/{id}", app.dataTankController.UpdateDataTank).Methods("PUT")
	app.router.HandleFunc("/data-tank/{id}", app.dataTankController.DeleteDataTank).Methods("DELETE")
	app.router.HandleFunc("/read-serial", app.dataTankController.ReadFromSerial).Methods("GET")
}

// handleShutdownSignal handles signals for graceful shutdown
func (app *App) handleShutdownSignal() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown signal
	<-stopChan

	// Allow time to finish last tasks (optional)
	gracefulShutdownTimeout := 20 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()

	// Shutdown the server with a timeout
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}

// Run runs the application
func (app *App) Run() {
	// Initialize Caching or Message Queue
	InitializeCache()

	// Register Middleware
	app.router.Use(MyLoggingMiddleware)

	// Initialize HTTP server
	app.server = &http.Server{
		Addr:         ":8080", // Change to the appropriate port
		Handler:      app.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Run HTTP server in a goroutine
	go func() {
		log.Printf("Server started on :8080\n") // Change to the appropriate port
		if err := app.server.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Handle signals for graceful shutdown
	app.handleShutdownSignal()
}
