package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex
	m := make(map[int]string)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			mu.Lock()
			m[i] = "a" // First conflicting access.
			mu.Unlock()
			wg.Done()
		}()
	}

	mu.Lock()
	m[2] = "b" // Second conflicting access.
	mu.Unlock()
	wg.Wait()
	for k, v := range m {
		fmt.Println(k, v)
	}
}
