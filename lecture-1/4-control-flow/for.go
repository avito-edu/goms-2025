package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		/*  */
	}

	/*  later in series */

	sl := []int{1, 2, 3}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	for range 3 {
		fmt.Println("Wake Up")
	}
}

func noWile() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}

		if i > 5 {
			break
		}
	}
}
