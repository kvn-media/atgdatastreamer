// internal/models/datatank.go

package models

import "time"

// DataTank adalah model untuk entitas DataTank
type DataTank struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Time           time.Time `json:"time"`
	Barel          int64     `json:"barel" validate:"required"`
	VolumeBarel    int       `json:"volume" validate:"required"`
	AveTemperature int       `json:"ave_temperature" validate:"required"`
	WaterDebit     int64     `json:"water" validate:"required"`
	TempProduct    int       `json:"temp_product" validate:"required"`
	Alarm          string    `json:"alarm"`
	// Tambahkan field lain sesuai kebutuhan
}
