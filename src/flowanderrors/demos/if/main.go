package main

import (
	"fmt"
)

func main() {

	x := 2

	// simple if statement
	if x == 2 {
		fmt.Println("x is 2")
	}

	// if with initialization statement
	if y := getY(); y == 5 {
		fmt.Println("y is 5")
	}

}

func getY() int {
	return 5
}
