package main

import "fmt"

// Простая функция без параметров и возвращаемых значений
func sayHello() {
	fmt.Println("Hello, World!")
}

// Функция с параметрами
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Функция с возвращаемым значением
func add(a int, b int) int {
	return a + b
}

// Сокращенная запись параметров одного типа
func multiply(x, y int) int {
	return x * y
}

func main() {
	sayHello()                // Hello, World!
	greet("Alice")            // Hello, Alice!
	sum := add(5, 3)          // 8
	product := multiply(4, 2) // 8

	fmt.Println("Sum:", sum, "Product:", product)
}
