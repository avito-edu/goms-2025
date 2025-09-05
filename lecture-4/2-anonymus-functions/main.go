package main

import (
	"fmt"
)

// 1) Счётчик
func Counter() func() int {
	x := 0
	return func() int { x++; return x }
}

// 2) Префиксер строк
func Prefixer(prefix string) func(string) string {
	return func(s string) string { return prefix + s }
}

// 3) Аккумулятор суммы
func Accumulator() func(int) int {
	sum := 0
	return func(v int) int { sum += v; return sum }
}

func main() {
	//f := func(name string) { fmt.Println("hi", name) }
	//f("Go")
	//
	//// Тип для читаемости
	//type StrOp func(string) string
	//trim := StrOp(func(s string) string { return strings.TrimSpace(s) })
	//
	//fmt.Println(trim("  hi  ")) // "hi"
	//
	//names := []string{"a", "b", "c", "d", "e", "f", "g"}
	//// Использование inline как компаратора (кейс‑инсensitive)
	//
	//slices.SortFunc(names, func(a, b string) int {
	//	return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	//})
	//
	//result := func(a, b int) int { return a * b }(2, 3)
	//fmt.Println(result)
	//
	//var port = func() int {
	//	if p := os.Getenv("PORT"); p != "" {
	//		n, _ := strconv.Atoi(p)
	//		return n
	//	}
	//	return 8080
	//}()
	//
	//fmt.Println(port)

	next := Counter()
	fmt.Println(next(), next()) // 1 2
	hello := Prefixer("Hello, ")
	fmt.Println(hello("Gophers")) // Hello, Gophers
	acc := Accumulator()
	fmt.Println(acc(3), acc(5)) // 3 8
}
