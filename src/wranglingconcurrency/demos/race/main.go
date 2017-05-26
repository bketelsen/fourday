package main

import "fmt"

func main() {

	//	var mu sync.Mutex
	c := make(chan bool)
	m := make(map[int]string)
	for i := 0; i < 1000; i++ {
		go func() {
			//		mu.Lock()
			m[i] = "a" // First conflicting access.
			//		mu.Unlock()
			c <- true
		}()
	}
	//	mu.Lock()
	m[2] = "b" // Second conflicting access.
	//	mu.Unlock()
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
