package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetUserInfo(t *testing.T) {
	// Создаём тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "Alice", "age": 30}`))
	}))
	defer server.Close() // Важно закрыть сервер после теста

	// Вызываем функцию, передавая URL тестового сервера
	result, err := GetUserInfo(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := `{"name": "Alice", "age": 30}`
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestGetUserInfo_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Имитируем долгий ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаём клиент с таймаутом
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	// Переопределяем клиент (в реальном коде лучше использовать интерфейсы)
	oldClient := http.DefaultClient
	defer func() { http.DefaultClient = oldClient }()
	http.DefaultClient = client

	_, err := GetUserInfo(server.URL)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}
