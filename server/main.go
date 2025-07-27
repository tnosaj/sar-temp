package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tnosaj/sar-temp/server/handler"
	"github.com/tnosaj/sar-temp/server/storage"
)

func main() {

	var apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("No API_KEY envvar set")
	}

	// Initialize storage (SQLite)
	store, err := storage.NewSQLiteStorage("data.db")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	h := &handler.Handler{Store: store}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/api/temperature", requireAPIKey(apiKey, h.PostTemperature))
	//http.HandleFunc("/api/dashboard", requireAPIKey(apiKey, h.GetDashboard))
	http.HandleFunc("/api/dashboard", h.GetDashboard)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func requireAPIKey(key string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("X-API-Key")
		if got != key {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
