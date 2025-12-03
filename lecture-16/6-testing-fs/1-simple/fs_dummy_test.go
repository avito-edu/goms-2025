package fs

import (
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	// Создаём временный файл
	err := os.WriteFile("test.txt", []byte("line1\nline2\nline3"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.txt") // Не всегда срабатывает при панике

	count, err := CountLines("test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}
