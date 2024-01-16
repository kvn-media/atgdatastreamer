package models

import "time"

// TankData represents data from the Automatic Tank Gauge system
type TankData struct {
	ID            int       `json:"id"`
	TankNumber    int       `json:"tank_number"`
	OilType       string    `json:"oil_type"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	StartVolume   float64   `json:"start_volume"`
	EndVolume     float64   `json:"end_volume"`
	DeliveryVolume float64   `json:"delivery_volume"`
	Temperature   float64   `json:"temperature"`
	// Add other relevant fields as needed
}