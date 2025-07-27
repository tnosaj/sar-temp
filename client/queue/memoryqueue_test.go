package queue

import (
	"testing"
	"time"
)

func TestMemoryQueue_AddAndPop(t *testing.T) {
	q := NewMemoryQueue()
	r := Reading{Timestamp: time.Now(), TemperatureC: 23.0}

	if err := q.Add(r); err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	items, err := q.PopAll()
	if err != nil {
		t.Fatalf("PopAll failed: %v", err)
	}
	if len(items) != 1 || items[0].TemperatureC != r.TemperatureC {
		t.Errorf("unexpected result: %+v", items)
	}
}
