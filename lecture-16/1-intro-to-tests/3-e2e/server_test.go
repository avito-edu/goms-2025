// server_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	HelloHandler(rr, req)

	if rr.Body.String() != "Hello, World!" {
		t.Errorf("Unexpected response: %s", rr.Body.String())
	}
}

func TestServer(t *testing.T) {
	// Запускаем сервер в тесте
	srv := httptest.NewServer(http.HandlerFunc(HelloHandler))
	defer srv.Close()

	// Делаем запрос к серверу
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверяем ответ
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
