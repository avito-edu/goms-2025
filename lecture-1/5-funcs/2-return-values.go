package main

import "fmt"

// Функция с несколькими возвращаемыми значениями
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

// Именованные возвращаемые значения
func power(base, exponent int) (result int) {
	result = 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return // автоматически возвращает result
}

// Функция с переменным числом аргументов
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	res, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", res) // 5
	}

	fmt.Println("2^5 =", power(2, 5)) // 32

	fmt.Println("Sum:", sum(1, 2, 3, 4))        // 10
	fmt.Println("Sum:", sum([]int{5, 6, 7}...)) // 18
}
