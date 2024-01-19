// internal/models/datatank.go

package models

// DataTank adalah model untuk entitas DataTank
type DataTank struct {
	ID          int     `json:"id"`
	Level       float64 `json:"level"`
	Temperature float64 `json:"temperature"`
	// Tambahkan field lain sesuai kebutuhan
}
