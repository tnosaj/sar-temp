package queue

import (
	"sync"
)

type MemoryQueue struct {
	mu       sync.Mutex
	readings []Reading
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		readings: make([]Reading, 0),
	}
}

func (q *MemoryQueue) Add(r Reading) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.readings = append(q.readings, r)
	return nil
}

func (q *MemoryQueue) PopAll() ([]Reading, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.readings) == 0 {
		return nil, nil
	}

	// Return a copy of the slice
	out := make([]Reading, len(q.readings))
	copy(out, q.readings)

	// Clear the queue
	q.readings = q.readings[:0]
	return out, nil
}
