package models

import "time"

// AlarmReport represents alarm report data
type AlarmReport struct {
	ID          int       `json:"id"`
	TankNumber  int       `json:"tank_number"`
	AlarmType   string    `json:"alarm_type"`
	AlarmTime   time.Time `json:"alarm_time"`
	Height      float64   `json:"height"`
	Temperature float64   `json:"temperature"`
	// Add other relevant fields as needed
}
