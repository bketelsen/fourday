package main

import (
	"fmt"
)

func main() {

	fmt.Println("Integer:")
	assertions(5)

	fmt.Println("String:")
	assertions("Gophers Rule")

	fmt.Println("Float:")
	assertions(3.14159)
}

func assertions(i interface{}) {

	switch v := i.(type) {
	case string:
		fmt.Println("String found", v)
	case int, int64:
		fmt.Println("Integer found", v)
	case float32, float64:
		fmt.Println("Float found", v)
	default:
		fmt.Println("unknown")
	}

}
