package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tnosaj/sar-temp/server/storage"
)

type mockStore struct {
	LastCalled string
	Data       []storage.TemperatureReading
}

func (m *mockStore) StoreTemperature(r storage.TemperatureReading) error {
	m.LastCalled = "StoreTemperature"
	m.Data = append(m.Data, r)
	return nil
}
func (m *mockStore) GetTodaysTemperatures(clientID string) ([]storage.TemperatureReading, error) {
	return m.Data, nil
}
func (m *mockStore) GetLastReadingTime(clientID string) (time.Time, error) {
	if len(m.Data) == 0 {
		return time.Time{}, nil
	}
	return m.Data[len(m.Data)-1].Timestamp, nil
}

func TestPostTemperature(t *testing.T) {
	h := &Handler{Store: &mockStore{}}
	payload := storage.TemperatureReading{
		ClientID:     "abc",
		Timestamp:    time.Now(),
		TemperatureC: 20.0,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)

	req := httptest.NewRequest("POST", "/api/temperature", b)
	w := httptest.NewRecorder()

	h.PostTemperature(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", w.Code)
	}
}

func TestGetDashboard(t *testing.T) {
	ms := &mockStore{}
	h := &Handler{Store: ms}

	req := httptest.NewRequest("GET", "/?client_id=test-client", nil)
	w := httptest.NewRecorder()
	h.GetDashboard(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
