package models

import "time"

// ShiftReport represents shift report data
type ShiftReport struct {
	ID          int       `json:"id"`
	ShiftNumber int       `json:"shift_number"`
	TankNumber  int       `json:"tank_number"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Succession  bool      `json:"succession"`
	StartVolume float64   `json:"start_volume"`
	EndVolume   float64   `json:"end_volume"`
	AddVolume   float64   `json:"add_volume"`
	SoldVolume  float64   `json:"sold_volume"`
	StartHeight float64   `json:"start_height"`
	EndHeight   float64   `json:"end_height"`
	// Add other relevant fields as needed
}
