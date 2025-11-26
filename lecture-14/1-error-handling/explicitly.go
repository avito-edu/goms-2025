package main

import (
	"fmt"
	"os"
)

func openFile() error {
	// Explicit handling (Go)
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("open config: %w", err)
	}
	defer file.Close()

	return nil
}

func main() {
	openFile()
}
