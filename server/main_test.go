package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAPIKey_Success(t *testing.T) {
	key := "testkey"
	protected := requireAPIKey(key, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-API-Key", "testkey")
	rr := httptest.NewRecorder()

	protected(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
}

func TestRequireAPIKey_Missing(t *testing.T) {
	key := "testkey"
	protected := requireAPIKey(key, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	protected(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401 Unauthorized, got %d", rr.Code)
	}
}

func TestRequireAPIKey_Invalid(t *testing.T) {
	key := "testkey"
	protected := requireAPIKey(key, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-API-Key", "wrongkey")
	rr := httptest.NewRecorder()

	protected(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401 Unauthorized, got %d", rr.Code)
	}
}
