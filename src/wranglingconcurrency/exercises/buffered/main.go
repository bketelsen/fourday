package main

import "fmt"

func main() {
	events := make(chan string, 1)

	events <- "Brian Logged In"

	e := <-events
	fmt.Println(e)
}
