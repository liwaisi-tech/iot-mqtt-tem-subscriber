package entities

import "time"

type ClimateDataEntity struct {
	ID          string
	DeviceID    string
	Temperature float64
	Humidity    float64
	CreatedAt   time.Time
}
