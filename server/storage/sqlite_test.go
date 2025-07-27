package storage

import (
	"os"
	"testing"
	"time"
)

func TestSQLiteStorage_CRUD(t *testing.T) {
	os.Remove("test.db")
	store, err := NewSQLiteStorage("test.db")
	if err != nil {
		t.Fatalf("failed to init storage: %v", err)
	}
	defer os.Remove("test.db")

	read := TemperatureReading{
		ClientID:     "test-client",
		Timestamp:    time.Now().UTC(),
		TemperatureC: 22.5,
	}

	if err := store.StoreTemperature(read); err != nil {
		t.Errorf("failed to store: %v", err)
	}

	today, err := store.GetTodaysTemperatures("test-client")
	if err != nil || len(today) == 0 {
		t.Errorf("expected reading, got %v", err)
	}

	last, err := store.GetLastReadingTime("test-client")
	if err != nil || last.IsZero() {
		t.Errorf("expected last timestamp, got %v", err)
	}
}
