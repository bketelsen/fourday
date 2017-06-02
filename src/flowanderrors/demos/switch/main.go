package main

import (
	"fmt"
)

func main() {

	sentence := "Now is the time for all good men to come to the aid of their country"
	fmt.Println(sentence)

	for _, b := range sentence {
		rot := rot13(byte(b))
		fmt.Printf(string(rot))
	}

}

// rot13 shows switch statements with complex case statements
func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}
