package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rec.Body.String())
	}
}

func TestGreetHandler(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"Valid name", "/greet?name=Alice", http.StatusOK, "Hello, Alice!"},
		{"Empty name", "/greet", http.StatusBadRequest, "Name is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rec := httptest.NewRecorder()

			GreetHandler(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
