package sensor

import (
	"testing"
)

func TestReadTemperature(t *testing.T) {
	temp := ReadTemperature()
	if temp < -50 || temp > 100 {
		t.Errorf("unrealistic temperature: %f", temp)
	}
}
