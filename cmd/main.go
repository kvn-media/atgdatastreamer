package main

import (
	"fmt"
	"os"

	"github.com/kvn-media/atgdatastreamer/internal/database"
	"github.com/kvn-media/atgdatastreamer/internal/serial"
)

func main() {
    dbPath := "synchub.db" // Sesuaikan dengan path database yang Anda inginkan
    _, err := database.InitDB(dbPath)
    if err != nil {
        fmt.Printf("Failed to initialize database: %v\n", err)
        os.Exit(1)
    }
    defer database.CloseDB()

    // Lanjutkan dengan inisialisasi komponen lain dan jalankan aplikasi
    // Inisialisasi SerialPort
    serialPort := serial.NewSerialPortImpl()
    err := serialPort.Connect("COM1", 9600) // Ganti dengan port dan baud rate yang sesuai
    if err != nil {
        fmt.Printf("Failed to connect to serial port: %v\n", err)
        os.Exit(1)
    }
    defer serialPort.Disconnect()

    // Inisialisasi usecase dengan SerialPort
    dataTankUsecase := usecase.NewDataTankUsecase(dataTankRepository, serialPort)

    // ... (jalankan aplikasi dan operasi lainnya)
}