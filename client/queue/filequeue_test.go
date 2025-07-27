package queue

import (
	"os"
	"testing"
	"time"
)

func TestFileQueue_AddAndPop(t *testing.T) {
	file := "test_queue.json"
	defer os.Remove(file)

	q := NewFileQueue(file)
	r := Reading{Timestamp: time.Now(), TemperatureC: 21.3}
	if err := q.Add(r); err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	readings, err := q.PopAll()
	if err != nil {
		t.Fatalf("PopAll failed: %v", err)
	}
	if len(readings) != 1 || readings[0].TemperatureC != r.TemperatureC {
		t.Errorf("unexpected readings: %+v", readings)
	}
}
