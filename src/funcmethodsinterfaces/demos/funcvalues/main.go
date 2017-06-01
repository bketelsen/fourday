package main

import (
	"fmt"
)

func main() {
	f := oneInput
	f("one")
	addSeparator("two", f)
}

// oneInput takes one input variable
// and returns no value
func oneInput(s string) {
	fmt.Println("oneInput", s)
}

func addSeparator(s string, fun func(string)) {
	fmt.Println("\n--------")
	fun(s)
}
