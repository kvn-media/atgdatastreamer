package models

import "time"

// ATGData represents the data collected from ATG
type ATGData struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	// Add other fields according to the data you collect from ATG
	Level     float64   `gorm:"not null"`
	Temperature float64 `gorm:"not null"`
	// Add more fields as needed
}
