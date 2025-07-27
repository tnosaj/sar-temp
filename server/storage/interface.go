package storage

import "time"

type TemperatureReading struct {
	ClientID     string    `json:"client_id"`
	Timestamp    time.Time `json:"timestamp"`
	TemperatureC float64   `json:"temperature_c"`
}

type Storage interface {
	StoreTemperature(reading TemperatureReading) error
	GetTodaysTemperatures(clientID string) ([]TemperatureReading, error)
	GetLastReadingTime(clientID string) (time.Time, error)
}
