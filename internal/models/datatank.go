// internal/models/datatank.go

package models

import "time"

// DataTank adalah model untuk entitas DataTank
type DataTank struct {
	ID             int       `json:"id"`
	Time           time.Time `json:"time"`
	Barel          int64     `json:"barel"`
	VolumeBarel    int       `json:"volume"`
	AveTemperature int       `json:"ave_temperature"`
	WaterDebit     int64     `json:"water"`
	TempProduct    int       `json:"temp_product"`
	Alarm          string    `json:"alarm"`
	// Tambahkan field lain sesuai kebutuhan
}
