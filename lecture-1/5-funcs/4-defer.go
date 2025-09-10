package main

import "fmt"

func main() {
	defer fmt.Println("Это выполнится третьим")
	defer fmt.Println("Это выполнится вторым")
	fmt.Println("Это выполнится первым")
}
