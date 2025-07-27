package queue

import "time"

type Reading struct {
	Timestamp    time.Time
	TemperatureC float64
}

type Queue interface {
	Add(Reading) error
	PopAll() ([]Reading, error)
}
