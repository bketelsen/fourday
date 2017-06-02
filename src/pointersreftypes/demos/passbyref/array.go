package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(&a)
	fmt.Println(a)
}
