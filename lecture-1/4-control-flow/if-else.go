package main

import "fmt"

func main() {
	flag := 1 == 23
	counter := 0

	if flag {
		counter += 1
	} else {
		counter -= 1
	}
	fmt.Println(flag, counter)
}
