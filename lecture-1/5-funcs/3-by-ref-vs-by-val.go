package main

import "fmt"

// Передача по значению (копируется)
func doubleValue(n int) {
	n *= 2
	fmt.Println("Inside doubleValue:", n) // 20
}

// Передача по ссылке (указатель)
func doublePointer(n *int) {
	*n *= 2
	fmt.Println("Inside doublePointer:", *n) // 20
}

func modifySlice(s []string) {
	s[0] = "Modified"
	fmt.Println("Inside modifySlice:", s) // [Modified two three]
}

func main() {
	// Передача по значению
	num := 10
	doubleValue(num)
	fmt.Println("After doubleValue:", num) // 10

	fmt.Println("~~~~~~~~~~")

	// Передача по ссылке
	doublePointer(&num)
	fmt.Println("After doublePointer:", num) // 20

	fmt.Println("~~~~~~~~~~")

	// Слайсы передаются по ссылке (на самом деле передается значение дескриптора слайса)
	words := []string{"one", "two", "three"}
	modifySlice(words)
	fmt.Println("After modifySlice:", words) // [Modified two three]

	fmt.Println("~~~~~~~~~~")

	// Для массивов нужно явно передавать указатель
	arr := [3]int{1, 2, 3}
	modifyArray(&arr)
	fmt.Println("After modifyArray:", arr) // [10 2 3]
}

func modifyArray(arr *[3]int) {
	arr[0] = 10
	fmt.Println("Inside modifyArray:", *arr) // [10 2 3]
}
