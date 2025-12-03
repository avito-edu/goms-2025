package concurrency

func ProcessAsync(data int, resultChan chan int) {
	go func() {
		resultChan <- data * 2
	}()
}

var counter int

func Increment() {
	counter++
}
