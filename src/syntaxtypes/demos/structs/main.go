// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type example set to its
	// zero value.
	var e1 example

	// Display the value.
	fmt.Printf("e1: %+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("e2: Flag", e2.flag)
	fmt.Println("e2: Counter", e2.counter)
	fmt.Println("e2: Pi", e2.pi)

}
