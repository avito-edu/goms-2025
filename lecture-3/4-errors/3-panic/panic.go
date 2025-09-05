package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero") // panic: division by zero
	}
	return a / b
}

func main() {
	fmt.Println(divide(10, 0))
}
