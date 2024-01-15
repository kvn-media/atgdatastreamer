package main

import (
	"database/sql"
	"log"
	"runtime"

	"github.com/kvn-media/atgdatastreamer/configs"
	"github.com/kvn-media/atgdatastreamer/controllers"
	"github.com/kvn-media/atgdatastreamer/managers"
	"github.com/kvn-media/atgdatastreamer/repositories"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use all CPU cores
}

func main() {
	// Load configuration
	config := configs.LoadConfig()
	if config == nil {
		log.Fatal("Error loading configuration")
	}

	// Connect to database
	db, err := sql.Open("sqlite3", "atg_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create ATG manager
	manager, err := managers.NewATGManager(config)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Close()

	// Create ATG repository
	repository := repositories.NewATGRepository(db)

	// Start collecting data
	go controllers.CollectATGData(manager, repository)

	// Block main thread to keep application running
	select {}
}
