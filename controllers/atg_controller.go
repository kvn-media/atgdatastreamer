package controllers

import (
	"fmt"
	"time"

	"github.com/kvn-media/atgdatastreamer/managers"
	"github.com/kvn-media/atgdatastreamer/repositories"
	"github.com/kvn-media/atgdatastreamer/services"
)

func CollectATGData(manager *managers.ATGManager, repository *repositories.ATGRepository) {
	// Continuously collect data from ATG
	for {
		// Read data from ATG
		data, err := services.ReadATGData(manager.Conn) // Access Conn through the manager instance
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Parse data into ATGData struct
		atgData, err := services.ParseATGData(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Save data to database
		err = repository.SaveATGData(atgData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Optionally, print data to console
		services.PrintATGData(atgData)

		// Do something else with the data, like send it to Mini PC 3

		// ...

		// Delay between readings to avoid flooding ATG
		time.Sleep(1 * time.Second)
	}
}
