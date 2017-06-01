package main

import (
	"fmt"
)

func main() {

	fmt.Println("Integer:")
	anything(5)

	fmt.Println("String:")
	anything("Gophers Rule")

	fmt.Println("Float:")
	anything(3.14159)
}

func anything(i interface{}) {
	fmt.Printf("%v\n\n", i)
}
