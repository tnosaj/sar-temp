package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tnosaj/sar-temp/client/queue"
	"github.com/tnosaj/sar-temp/client/sensor"
)

const (
	clientID = "raspi-001"
	interval = time.Minute
	server   = "http://127.0.0.1:8080/api/temperature"
)

func main() {
	var apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("No API_KEY envvar set")
	}
	q := queue.NewMemoryQueue()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		temp := sensor.ReadTemperature()
		q.Add(queue.Reading{
			Timestamp:    time.Now().UTC(),
			TemperatureC: temp,
		})

		trySend(q, apiKey)
		<-ticker.C
	}
}

func trySend(q queue.Queue, apiKey string) {
	readings, err := q.PopAll()
	if err != nil || len(readings) == 0 {
		return
	}
	for _, r := range readings {
		payload := map[string]interface{}{
			"client_id":     clientID,
			"timestamp":     r.Timestamp,
			"temperature_c": r.TemperatureC,
		}
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(payload)

		req, _ := http.NewRequest("POST", server, buf)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", apiKey)

		resp, err := http.DefaultClient.Do(req)

		if err != nil || resp.StatusCode >= 300 {
			log.Println("Failed to send data, re-queuing")
			q.Add(r) // naive retry
		}
	}
}
