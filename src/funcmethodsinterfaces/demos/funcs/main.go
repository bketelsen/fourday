package main

import (
	"errors"
	"fmt"
)

func main() {
	printSeparator()
	oneInput("one")

	printSeparator()
	twoInputs("one", "two")

	printSeparator()
	variadic("one")

	printSeparator()
	variadic("multiple", "input", "values")

	printSeparator()
	timesTwo(2) // no output!  WHY?

	printSeparator()
	four := timesTwo(2)
	fmt.Println("timesTwo", four)

	printSeparator()
	x, err := upToTen(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

	printSeparator()
	x, err = upToTen(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

	printSeparator()
	// that's a lot of typing above, let's simplify it
	if x, err = upToTen(9); err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

}

func printSeparator() {
	fmt.Println("\n--------------------")
}

// oneInput takes one input variable
// and returns no value
func oneInput(s string) {
	fmt.Println("oneInput", s)
}

// twoInputs takes two input variables
// and returns no value
func twoInputs(s, t string) {
	fmt.Println("twoInputs", s)
	fmt.Println("twoInputs", t)
}

// variadic arguments takes zero or more inputs
// which are presented to the function as a slice.
// note the three dots in front of the type, that's
// the notation for a variadic parameter.
func variadic(s ...string) {
	fmt.Printf("variadic received %d inputs:\n", len(s))
	for idx, val := range s {
		fmt.Printf("\tIndex %d\t Value %s\n", idx, val)
	}

}

// timesTwo takes an integer input
// and returns that integer multiplied by two
func timesTwo(i int) int {
	return i * 2
}

// upToTen takes an integer input
// and adds 5 to it.  It returns the result
// if the result is less than 10, or it returns
// an error.
func upToTen(i int) (int, error) {
	x := i + 5
	if x > 9 {
		return i, errors.New("number too big")
	}
	return x, nil

}
