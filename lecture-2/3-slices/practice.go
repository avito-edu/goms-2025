package main

import "fmt"

func main() {
	// Create an empty slice (length 0, capacity 0)
	var nums []int
	fmt.Printf("Initial: len=%d, cap=%d\n", len(nums), cap(nums)) // len=0, cap=0

	// Add 100 elements in a loop
	for i := 0; i < 100; i++ {
		nums = append(nums, i)

		// Print length & capacity when it changes (optional)
		// if i%10 == 0 {
		//	fmt.Printf("After %3d appends: len=%3d, cap=%3d\n", i+1, len(nums), cap(nums))
		// }
	}

	// Final length and capacity
	fmt.Printf("Final: len=%d, cap=%d\n", len(nums), cap(nums)) // len=100, cap=128
}
