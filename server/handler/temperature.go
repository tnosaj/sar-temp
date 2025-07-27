package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tnosaj/sar-temp/server/storage"
)

type Handler struct {
	Store storage.Storage
}

func (h *Handler) PostTemperature(w http.ResponseWriter, r *http.Request) {
	var reading storage.TemperatureReading
	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if reading.Timestamp.IsZero() {
		reading.Timestamp = time.Now()
	}
	if err := h.Store.StoreTemperature(reading); err != nil {
		http.Error(w, "store error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// In handler/temperature.go
func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		http.Error(w, "client_id required", http.StatusBadRequest)
		return
	}
	readings, err := h.Store.GetTodaysTemperatures(clientID)
	if err != nil {
		http.Error(w, "failed to get readings", http.StatusInternalServerError)
		return
	}
	last, err := h.Store.GetLastReadingTime(clientID)
	if err != nil {
		last = time.Time{} // fallback to zero value
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"readings":  readings,
		"last_seen": last,
	})
}
