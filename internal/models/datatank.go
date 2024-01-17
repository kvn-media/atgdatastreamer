package models

// DataTank adalah model untuk entitas DataTank
type DataTank struct {
    ID     int    `json:"id"`
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
    Field3 float64 `json:"field3"`
    // Tambahkan field lain sesuai kebutuhan
}