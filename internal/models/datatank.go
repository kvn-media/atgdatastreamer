// internal/models/datatank.go

package models

// DataTank adalah model untuk entitas DataTank
type DataTank struct {
	ID        int     `json:"id"`
    Name      string  `json:"name"`
    Capacity  int     `json:"capacity"`
    Temperature float64 `json:"temperature"`
	// Tambahkan field lain sesuai kebutuhan
}
