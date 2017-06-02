package main

import (
	"fmt"
)

func main() {

	// for with init, condition, post
	// just like c, java, etc
	var x int
	for x = 1; x < 10; x++ {
		fmt.Println("x:", x)
	}

	fmt.Println("*******")

	// use a short declaration to make it prettier
	for y := 0; y < 9; y++ {
		fmt.Println("y:", y)
	}

	fmt.Println("*******")

	z := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arrays, slices, strings, maps and channels support
	// iterating with range which
	// automatically manage the loop
	for k, v := range z {
		fmt.Printf("index %d = %d\n", k, v)
	}

	fmt.Println("*******")

	// ranging over strings is unicode aware and returns
	// the byte position and unicode char
	for pos, char := range "日本語" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

}
