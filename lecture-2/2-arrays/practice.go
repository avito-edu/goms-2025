package main

import "fmt"

func main() {
	// Create an array of 5 integers
	numbers := [5]int{10, 20, 30, 40, 50}

	// Print array length
	fmt.Println("Array length:", len(numbers)) // Output: 5

	// Print first element (index 0)
	fmt.Println("First element:", numbers[0]) // Output: 10

	// Print last element (index len-1)
	lastIndex := len(numbers) - 1
	fmt.Println("Last element:", numbers[lastIndex]) // Output: 50
}
