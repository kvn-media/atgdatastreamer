// internal/application/cache.go

package application

import (
	"fmt"
	"time"

	"github.com/kvn-media/atgdatastreamer/internal/models"
	"github.com/patrickmn/go-cache"
)

var globalCache *cache.Cache

// InitializeCache initializes the cache
func InitializeCache() {
	// Initialize the global cache with a default expiration time of 5 minutes
	globalCache = cache.New(5*time.Minute, 10*time.Minute)

	fmt.Println("Cache initialized")
}

// Example of using the cache
func ExampleUsingCache() {
	// Assuming you have a tank with ID 123 and some associated data
	var (
		tankID = 1
		timeStamp = time.Now
		barel = 1000
		volumeBar = -20
		aveTemperature = 20
		waterBar = 0.1
		tempProduct = 30
		notification = "Tidak ada"
	) 
	
	key := fmt.Sprintf("tankData:%d", tankID)

	// Create a sample TankData instance
	value := &models.DataTank{
		ID:          tankID,
		Time: timeStamp(),
		Barel: int64(barel),
		VolumeBarel: volumeBar,
		AveTemperature: aveTemperature,
		WaterDebit: int64(waterBar),
		TempProduct: tempProduct,
		Alarm: notification,
	}

	// Store the TankData instance in the cache
	globalCache.Set(key, value, cache.DefaultExpiration)

	// Retrieve the TankData instance from the cache
	cachedValue, found := globalCache.Get(key)
	if found {
		fmt.Printf("Cached data for tank with ID '%d': %v\n", tankID, cachedValue)
	} else {
		fmt.Printf("Data for tank with ID '%d' not found in the cache\n", tankID)
	}
}