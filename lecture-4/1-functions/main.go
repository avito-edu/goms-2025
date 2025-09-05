package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("zero")
	}
	return a / b, nil
}

func bounds(n int) (min, max int) {
	min, max = 0, n
	return // возвращает min, max
}

func add(a, b int) int { return a + b }

func choose(op string) func(int, int) int {
	if op == "+" {
		return func(a, b int) int { return a + b }
	}
	return func(a, b int) int { return a - b }
}

func apply(a, b int, f func(int, int) int) int { return f(a, b) }

func main() {
	var op func(int, int) int = add
	fmt.Println(op(2, 3)) // 5

	fmt.Println(apply(2, 3, func(x, y int) int { return x * y }))
}
