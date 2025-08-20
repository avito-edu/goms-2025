package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// createBinary()

	// Открытие бинарного файла для чтения
	file, err := os.Open("./lecture-3/1-read-file/binary/person.gob")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var person Person
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding data:", err)
	} else {
		fmt.Println("Read data:", person)
	}
}

func createBinary() {
	person := Person{
		Name: "Denis Zakharov",
		Age:  25,
	}

	file, err := os.Create("./lecture-3/1-read-file/binary/person.gob")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding data:", err)
	} else {
		fmt.Println("Data successfully written to person.gob")
	}
}
