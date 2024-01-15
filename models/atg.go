package models

import (
    "time"
)

type ATGData struct {
    Type string
    ID int
    Value float64
    Timestamp time.Time
}