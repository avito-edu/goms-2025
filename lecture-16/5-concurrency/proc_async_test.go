package concurrency

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestProcessAsync_Bad(t *testing.T) {
	resultChan := make(chan int)
	ProcessAsync(5, resultChan)

	// Тест завершится, не дожидаясь результата!
	if result := <-resultChan; result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

func TestProcessAsync_Good(t *testing.T) {
	resultChan := make(chan int, 1) // Буферизованный канал
	ProcessAsync(5, resultChan)

	select {
	case result := <-resultChan:
		if result != 10 {
			t.Errorf("Expected 10, got %d", result)
		}
	case <-time.After(1 * time.Second): // Таймаут на случай зависания
		t.Error("Timeout waiting for result")
	}
}

func TestIncrement_Race(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go Increment()
	}

	time.Sleep(1 * time.Second)    // Наивная попытка дождаться завершения
	t.Logf("Counter: %d", counter) // Результат непредсказуем!
}

func ProcessBatch(data []int, resultChan chan int) {
	var wg sync.WaitGroup
	for _, num := range data {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultChan <- n * 2
		}(num)
	}
	wg.Wait()
	close(resultChan)
}

func TestProcessBatch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int, len(data))

	ProcessBatch(data, resultChan)

	var results []int
	for res := range resultChan {
		results = append(results, res)
	}

	expected := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}

type Cache struct {
	mu    sync.Mutex
	items map[string]string
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.items[key]
	return val, ok
}

func TestCache_ConcurrentAccess(t *testing.T) {
	cache := &Cache{items: make(map[string]string)}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			cache.Set(key, "value")
		}(i)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := cache.Get(key); !ok || val != "value" {
			t.Errorf("Key %s not found or incorrect value", key)
		}
	}
}
