package fs

import (
	"testing"

	"github.com/spf13/afero"
)

func TestCountLinesWithAfero(t *testing.T) {
	// Создаём мок файловой системы
	fs := afero.NewMemMapFs()

	// Создаём файл в памяти
	err := afero.WriteFile(fs, "test.txt", []byte("line1\nline2\nline3"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	count, err := CountLines(fs, "test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}

func TestCountLines_FileNotExists(t *testing.T) {
	fs := afero.NewMemMapFs()
	_, err := CountLines(fs, "nonexistent.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}
