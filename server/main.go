package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	// Serve static frontend (e.g., /static/index.html)
	staticDir := "static"
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		if r.URL.Query().Has("client_id") || accept == "application/json" {
			requireAPIKey(apiKey, h.GetDashboard)(w, r)
			return
		}
		// Serve HTML dashboard page
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})

	http.HandleFunc("/api/temperature", requireAPIKey(apiKey, h.PostTemperature))

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
