package main

import "fmt"

func main() {
	events := make(chan string)

	go func() {
		events <- "Brian Logged In"
	}()

	e := <-events
	fmt.Println(e)
}
