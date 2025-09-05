package main

import (
	"fmt"
	"log"
	"os"
)

const (
	// filePath — to the curr dir the file to be created/opened/deleted
	filePath = "./lecture-3/2-opening-modes"
	fileName = "dummy.txt"
)

func main() {
	fp := fmt.Sprintf("%s/%s", filePath, fileName)

	// Create or truncate for writing only (O_WRONLY|O_CREATE|O_TRUNC)
	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file.Close()

	file.WriteString("This is the first line.\n")
	file.Write([]byte("This is a byte slice.\n"))

	// Open for reading and writing (O_RDWR)
	file, err = os.OpenFile(fp, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file.Close()

	// Open for appending (O_WRONLY|O_APPEND)
	file, err = os.OpenFile(fp, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file.Close()

	err = os.Remove(fp)
	// Handle any potential errors during the removal process
	if err != nil {
		log.Fatalf("Error deleting file %s: %v", filePath, err)
	}
}
