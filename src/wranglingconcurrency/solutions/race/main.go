package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	var mu sync.Mutex
	go func() {
		mu.Lock()
		defer mu.Unlock()
		m["1"] = "a"
		c <- true
	}()
	mu.Lock()
	m["2"] = "b"
	mu.Unlock()
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
