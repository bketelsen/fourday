package main

import (
	"fmt"
)

func main() {

	x := 5
	y := 6

	gt := greaterThan(x, y)
	fmt.Printf("%d > %d = %t", x, y, gt)
}

// greaterThan returns true if the first
// parameter is greater than the second.
func greaterThan(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
