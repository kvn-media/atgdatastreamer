package main

import (
	"github.com/kvn-media/atgdatastreamer/internal/application"
	"github.com/kvn-media/atgdatastreamer/internal/controllers"
)

func main() {
	// Create an instance of DataTankController
    dataTankController := new(controllers.DataTankController)

    // Create an instance of App and pass DataTankController to it
    app := application.NewApp(dataTankController)

    // Initialize the application
    app.Initialize()

    // Run the application
    app.Run()
}
