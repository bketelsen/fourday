package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	defer fmt.Println("First Defer")

	defer fmt.Println("Second Defer")

	defer fmt.Println("Third Defer")

	panic("Boom!")
}
