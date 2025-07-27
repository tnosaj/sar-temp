package sensor

import "math/rand"

func ReadTemperature() float64 {
	// Replace with GPIO sensor read
	return 20.0 + rand.Float64()*5.0
}
