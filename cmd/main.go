// main.go
package main

import (
	"github.com/kvn-media/atgdatastreamer/internal/application"
)

func main() {
	app := application.NewApp()
	// app.initializeRoutes()
	app.Run()
}