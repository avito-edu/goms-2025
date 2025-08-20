package main

import (
	"fmt"
	"log"
	"os"
)

func simpleReadFile() {
	f, err := os.ReadFile("./lecture-3/README.md")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(f))
}

func main() {

}
