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
}

func NewApp(dataTankController controllers.DataTankController) *App {
	return &App{
		dataTankController: dataTankController,
	}
}

func (app *App) Initialize() {
	// Load configuration from an external file
	config, err := configs.LoadConfig("configs/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database
	app.db, err = database.InitDB(config.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.CloseDB(app.db)

	// Initialize repository and usecase
	dataTankRepository := repository.NewDataTankRepository(app.db)
	serialPort := serial.NewSerialPortImpl()
	err = serialPort.Connect(config.SerialPortName, config.SerialPortBaud)
	if err != nil {
		log.Fatalf("Failed to connect to the serial port: %v", err)
	}
	defer serialPort.Disconnect()

	// Initialize HTTPS Delivery
	httpsDelivery := delivery.NewHttpsDelivery(config.HTTPSEndpoint)

	dataTankUsecase := usecase.NewDataTankUsecase(dataTankRepository, serialPort, httpsDelivery)
	app.dataTankController = controllers.NewDataTankController(dataTankUsecase)

	// Initialize the router
	app.router = mux.NewRouter()
	app.initializeRoutes()
}

// initializeRoutes menambahkan rute-rute ke router
func (app *App) initializeRoutes() {
	app.router.HandleFunc("/data-tank", app.dataTankController.CreateDataTank).Methods("POST")
	app.router.HandleFunc("/data-tank", app.dataTankController.GetDataTanks).Methods("GET")
	app.router.HandleFunc("/data-tank/{id}", app.dataTankController.UpdateDataTank).Methods("PUT")
	app.router.HandleFunc("/data-tank/{id}", app.dataTankController.DeleteDataTank).Methods("DELETE")
	app.router.HandleFunc("/read-serial", app.dataTankController.ReadFromSerial).Methods("GET")
}

// handleShutdownSignal menangani sinyal untuk graceful shutdown
func (app *App) handleShutdownSignal() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Menunggu sinyal shutdown
	<-stopChan

	// Memberikan waktu untuk menyelesaikan tugas terakhir (opsional)
	gracefulShutdownTimeout := 20 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()

	// Memberhentikan server dengan timeout
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}

// Run menjalankan aplikasi
func (app *App) Run() {
	// Migrasi Database
	err := database.PerformDatabaseMigration(app.db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Inisialisasi Caching atau Message Queue
	InitializeCache()

	// Registrasi Middleware
	app.router.Use(MyLoggingMiddleware)

	// Inisialisasi server HTTP
	app.server = &http.Server{
		Addr:         ":8080", // Ganti dengan port yang sesuai
		Handler:      app.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Jalankan server HTTP dalam goroutine
	go func() {
		log.Printf("Server started on :8080\n") // Ganti dengan port yang sesuai
		if err := app.server.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Menangani sinyal untuk graceful shutdown
	app.handleShutdownSignal()
}
