package main

import "testing"
import "net/http"
import "net/http/httptest"

func TestMain(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"localhost:8080/metrics",
		nil,
	)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %d", rec.Code)
	}
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/weather",
		nil,
	)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected code 200, got %d", rec.Code)
	}
}
